package http_middleware

import (
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
)

func HTTPServiceUserFlowLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := util.GetUser(c)
		if user == nil || err != nil {
			c.Next()
			return
		}
		// 对用户限流,限定用户名称
		userQPS := user.Qps
		if userQPS > 0 {
			clientLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+user.AppID, int(userQPS))
			if !clientLimiter.Limit.Allow() {
				util.RspError(c, util.CodeClientRateLimit, nil)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
