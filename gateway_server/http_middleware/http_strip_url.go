package http_middleware

import (
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func HTTPServiceStripUrl() gin.HandlerFunc {
	// example
	// rule: /test_http_service/abc
	// after strip: /abc
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		if detail.HttpRule.RuleType == global.PathType && detail.HttpRule.NeedStripUri == 1 {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, detail.HttpRule.Rule, "", 1)
		}
		c.Next()
	}
}
