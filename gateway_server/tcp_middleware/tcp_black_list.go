package tcp_middleware

import (
	"fmt"
	"gateway_server/cache/model"
	"gateway_server/tcp_middleware/tin"
	"strings"
)

func TcpBlackListMiddleware() tin.TinHandler {
	return func(c *tin.TinContext) {
		s := c.Get("service")
		if s == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		detail := s.(*model.ServiceDetail)
		white := detail.AccessControl.WhiteList
		split := strings.Split(c.Conn.RemoteAddr().String(), ":")
		clientIp := ""
		if len(split) == 2 {
			clientIp = split[0]
		}
		if detail.AccessControl.OpenAuth == 1 && len(white) == 0 && len(detail.AccessControl.BlackList) > 0 {
			if InSlice(strings.Split(detail.AccessControl.BlackList, ","), clientIp) {
				c.Conn.Write([]byte(fmt.Sprintf("%s in blacklist", clientIp)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
