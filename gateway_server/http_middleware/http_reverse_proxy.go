package http_middleware

import (
	"gateway_server/reverse_proxy"
	"gateway_server/reverse_proxy/load_balance"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"github.com/pretty66/websocketproxy"
	"net/http"
)

func HTTPReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		lb := load_balance.GetLoadBalanceByService(detail)
		if detail.HttpRule.NeedWebsocket == 1 {
			proxy, err := websocketproxy.NewProxy(lb.Get("")+c.Request.URL.Path, func(r *http.Request) error {
				return nil
			})
			if err != nil {
				c.Abort()
				return
			}
			proxy.ServeHTTP(c.Writer, c.Request)
			c.Next()
		} else {
			trans := load_balance.GetTransport(detail)
			proxy := reverse_proxy.NewHttpReverseProxyWithLoadBalance(c, lb, trans)
			proxy.ServeHTTP(c.Writer, c.Request)
		}
	}
}
