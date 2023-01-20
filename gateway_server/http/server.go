package http

import (
	"fmt"
	"gateway_server/global"
	"gateway_server/http_middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var HTTPServer http.Server

func StartHTTPServer() {
	conf := global.ProxyFullConfig.HTTPConfig
	engine := gin.Default()
	engine.Use(
		http_middleware.HTTPAccessMode(),             // 通过请求判断需要处理的service
		http_middleware.HTTPServiceFlowCount(),       // 流量统计
		http_middleware.HTTPServiceFlowLimit(),       // 流量限制
		http_middleware.JWTAuth(),                    // jwt验证用户身份
		http_middleware.HTTPServiceUserFlowCount(),   // 用户流量统计
		http_middleware.HTTPServiceUserFlowLimit(),   // 用户限流
		http_middleware.HTTPServiceWhiteList(),       // 白名单
		http_middleware.HTTPServiceBlackList(),       // 黑名单
		http_middleware.HTTPServiceUrlRewrite(),      // 根据配置的正则表达式对path进行重写
		http_middleware.HTTPServiceHeaderTransform(), // 根据配置进行header头的替换
		http_middleware.HTTPServiceStripUrl(),
		http_middleware.HTTPReverseProxy(),
	)
	HTTPServer = http.Server{
		Addr:           fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler:        engine.Handler(),
		ReadTimeout:    time.Duration(conf.ReadTimeout),
		WriteTimeout:   time.Duration(conf.WriteTimeout),
		MaxHeaderBytes: conf.MaxHeaderBytes,
	}
	zap.S().Info("[Navi Gateway] start http proxy server!")
	if err := HTTPServer.ListenAndServe(); err != nil {
		zap.S().Errorf("[Navi Gateway] http server start error,ip:%s", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
		return
	}
}

func CloseHTTPServer() {
	HTTPServer.Close()
}
