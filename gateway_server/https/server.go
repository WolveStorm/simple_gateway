package https

import (
	"fmt"
	"gateway_server/global"
	"gateway_server/http_middleware"
	"gateway_server/https/ca"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

////CA私钥
//openssl genrsa -out ca.key 2048
////CA数据证书
//openssl req -x509 -new -nodes -key ca.key -subj "/CN=example1.com" -days 5000 -out ca.crt
////服务器私钥（默认由CA签发）
//openssl genrsa -out server.key 2048
////服务器证书签名请求：Certificate Sign Request，简称csr（example1.com代表你的域名）
//openssl req -new -key server.key -subj "/CN=example1.com" -out server.csr
////上面2个文件生成服务器证书（days代表有效期）
//openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

var HTTPSServer *http.Server

func StartHTTPSServer() {
	conf := global.ProxyFullConfig.HTTPSConfig
	engine := gin.Default()
	engine.Use(
		http_middleware.HTTPSAccessMode(),            // 通过请求判断需要处理的service
		http_middleware.HTTPServiceFlowCount(),       // 流量统计
		http_middleware.HTTPServiceFlowLimit(),       // 流量限制
		http_middleware.JWTAuth(),                    // jwt验证用户身份
		http_middleware.HTTPServiceUserFlowCount(),   // 用户流量统计
		http_middleware.HTTPServiceUserFlowLimit(),   // 用户限流
		http_middleware.HTTPServiceWhiteList(),       // 白名单
		http_middleware.HTTPServiceBlackList(),       // 黑名单
		http_middleware.HTTPServiceUrlRewrite(),      // 根据配置的正则表达式对path进行重写
		http_middleware.HTTPServiceHeaderTransform(), // 根据配置进行header头的替换
		http_middleware.HTTPServiceStripUrl(),        //删除path前导的rule
		http_middleware.HTTPReverseProxy(),
	)
	zap.S().Info("[Navi Gateway] start https proxy server!")
	if err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", conf.Host, conf.Port), ca.Path("server.crt"), ca.Path("server.key"), engine.Handler()); err != nil {
		zap.S().Errorf("[Navi Gateway] http server start error,ip:%s", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
		return
	}
}

func CloseHTTPSServer() {
	HTTPSServer.Close()
}
