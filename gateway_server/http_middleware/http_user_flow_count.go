package http_middleware

import (
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"sync/atomic"
)

func HTTPServiceUserFlowCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 没有找到用户就不计数了，因为有可能是游客，并且服务没有open auth
		user, err := util.GetUser(c)
		if user == nil || err != nil {
			c.Next()
			return
		}
		counter := cache.FlowManager.GetFlowCounter(global.UserFlowLimit + user.AppID)
		atomic.AddInt64(&counter.AddTicket, 1)
		// 用户日请求量限制
		if user.Qpd > 0 {
			if user.Qpd <= counter.LastTicket {
				util.RspError(c, util.CodeUserRateLimit, nil)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
