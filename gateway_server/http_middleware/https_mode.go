package http_middleware

import (
	"encoding/json"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"github.com/gin-gonic/gin"
	"strings"
)

// 通过请求映射到后台的配置
func HTTPSAccessMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := cache.GetAllHTTPService()
		var req = c.Request
		var choosed *model.ServiceDetail
		for _, v := range service {
			if v.HttpRule.NeedHttps == 0 {
				continue
			}
			// 得到domain并且比较，则说明映射到了
			if v.HttpRule.RuleType == global.DomainType && req.Host[:strings.Index(req.Host, ":")] == v.HttpRule.Rule {
				choosed = v
				break
			}
			// 若path和配置的path有前缀重复，则说明映射到了
			if v.HttpRule.RuleType == global.PathType && strings.Contains(req.URL.Path, v.HttpRule.Rule) {
				choosed = v
				break
			}
		}
		marshal, _ := json.Marshal(choosed)
		c.Set(global.ServiceId, choosed.Id)
		c.Set(global.ServiceDetail, string(marshal))
		c.Next()
	}
}
