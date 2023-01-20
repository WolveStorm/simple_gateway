package http_middleware

import (
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func HTTPServiceUrlRewrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		// example
		// pattern like : ^/test_http_service/abb/(.*) /test_http_service/bba/$1
		// request like : /test_http_service/abb/1/2/3/4
		// rewrite like : /test_http_service/bba/1/2/3/4
		if detail.HttpRule.UrlRewrite != "" {
			patterns := strings.Split(detail.HttpRule.UrlRewrite, "\n")
			for _, p := range patterns {
				split := strings.Split(p, " ")
				if len(split) != 2 {
					continue
				}
				match := split[0]
				rewrite := split[1]
				cp, err := regexp.Compile(match)
				if err != nil {
					continue
				}
				path := cp.ReplaceAll([]byte(c.Request.URL.Path), []byte(rewrite))
				c.Request.URL.Path = string(path)
			}
		}
		c.Next()
	}
}
