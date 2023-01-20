package http_middleware

import (
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HTTPServiceUpgrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		// 对协议做一个升级，到时候返回websocket数据就通过这个conn返回即可。
		if detail.HttpRule.NeedWebsocket == 1 {
			wsConn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
			c.Set("websocket", wsConn)
			c.Next()
		}
		c.Next()
	}
}
