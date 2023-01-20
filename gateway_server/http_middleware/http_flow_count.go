package http_middleware

import (
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"sync/atomic"
)

func HTTPServiceFlowCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		counter := cache.FlowManager.GetFlowCounter(detail.ServiceInfo.ServiceName)
		atomic.AddInt64(&counter.AddTicket, 1)
		totalCounter := cache.FlowManager.GetFlowCounter(global.TotalKey)
		atomic.AddInt64(&totalCounter.AddTicket, 1)
		c.Next()
	}
}
