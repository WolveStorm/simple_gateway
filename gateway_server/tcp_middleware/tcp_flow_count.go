package tcp_middleware

import (
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"gateway_server/tcp_middleware/tin"
	"sync/atomic"
)

func TcpFlowCountMiddleware() tin.TinHandler {
	return func(c *tin.TinContext) {
		s := c.Get("service")
		if s == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		detail := s.(*model.ServiceDetail)
		counter := cache.FlowManager.GetFlowCounter(detail.ServiceInfo.ServiceName)
		atomic.AddInt64(&counter.AddTicket, 1)
		totalCounter := cache.FlowManager.GetFlowCounter(global.TotalKey)
		atomic.AddInt64(&totalCounter.AddTicket, 1)
		c.Next()
	}
}
