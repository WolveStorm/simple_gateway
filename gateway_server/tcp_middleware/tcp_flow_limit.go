package tcp_middleware

import (
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"gateway_server/tcp_middleware/tin"
	"strings"
)

func TcpFlowLimitMiddleware() tin.TinHandler {
	return func(c *tin.TinContext) {
		s := c.Get("service")
		if s == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		detail := s.(*model.ServiceDetail)
		//得到流量限制
		qps := float64(detail.AccessControl.ServiceFlowLimit)
		if qps > 0 {
			limiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName, int(qps))

			//对服务器限流
			if !limiter.Limit.Allow() {
				c.Conn.Write([]byte("flow limit!"))
				c.Abort()
				return
			}
		}
		split := strings.Split(c.Conn.RemoteAddr().String(), ":")
		clientIp := ""
		if len(split) == 2 {
			clientIp = split[0]
		}

		//对客户端ip限流
		qps1 := float64(detail.AccessControl.ClientIPFlowLimit)
		if qps1 > 0 {
			clientLimit := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName+":"+clientIp, int(qps1))
			//限流
			if !clientLimit.Limit.Allow() {
				c.Conn.Write([]byte("您已被限流，请稍等"))
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
