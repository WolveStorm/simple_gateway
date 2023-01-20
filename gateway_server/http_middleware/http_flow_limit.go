package http_middleware

import (
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
)

// 进行客户端和服务器的限流
func HTTPServiceFlowLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		// 对服务器限流，限定服务器
		serverQps := detail.AccessControl.ServiceFlowLimit
		if serverQps > 0 {
			serverLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName, serverQps)
			if !serverLimiter.Limit.Allow() {
				util.RspError(c, util.CodeServerRateLimit, nil)
				c.Abort()
				return
			}
		}

		// 对客户端限流,限定客户端IP
		clientQps := detail.AccessControl.ClientIPFlowLimit
		if clientQps > 0 {
			clientLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName+c.ClientIP(), clientQps)
			if !clientLimiter.Limit.Allow() {
				util.RspError(c, util.CodeClientRateLimit, nil)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
