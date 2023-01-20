package http_middleware

import (
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

// notice: 开启验证则会验证黑白名单那,白名单优先于黑名单，所以在中间件链路需要靠前
func HTTPServiceWhiteList() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		if detail.AccessControl.WhiteList != "" && detail.AccessControl.OpenAuth == 1 {
			whiteIps := strings.Split(detail.AccessControl.WhiteList, "\n")
			var match bool
			for _, ip := range whiteIps {
				if ip == c.ClientIP() {
					match = true
					break
				}
			}
			if !match {
				util.RspError(c, util.CodeNotInWhiteIPS, nil)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
