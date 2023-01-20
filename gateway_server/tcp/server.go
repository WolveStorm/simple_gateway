package tcp

import (
	"context"
	"fmt"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"gateway_server/reverse_proxy/load_balance"
	"gateway_server/tcp/server"
	"gateway_server/tcp_middleware"
	"gateway_server/tcp_middleware/tin"
	"go.uber.org/zap"
)

var tcpList = []*server.TCPServer{}

func StartTCPServer() {
	service := cache.GetAllTCPService()
	for _, v := range service {
		go func(detail *model.ServiceDetail) {
			router := tin.NewTinRouter()
			group := router.Group("/")
			group.Use(
				tcp_middleware.TcpFlowCountMiddleware(),
				tcp_middleware.TcpFlowLimitMiddleware(),
				tcp_middleware.TcpWhiteListMiddleware(),
				tcp_middleware.TcpBlackListMiddleware(),
			)
			lb := load_balance.GetLoadBalanceByService(detail)
			handler := tin.NewTinSliceRouterHandler(func(c *tin.TinContext) server.TCPHandler {
				// 核心方法只做反向代理
				return tcp_middleware.TCPReverseProxyWithLoadBalance(c, lb)
			}, router)
			ctx := context.Background()
			ctx = context.WithValue(ctx, "service", detail)
			s, err := server.ListenAndServe(fmt.Sprintf("%s:%d", global.ProxyFullConfig.TCPConfig.Host, detail.TcpRule.Port), handler, ctx)
			if err != nil {
				zap.S().Errorf("[Navi Gateway] start tcp proxy server error,port:%s", global.ProxyFullConfig.TCPConfig.Host)
			}
			tcpList = append(tcpList, s)
		}(v)
	}
}

func CloseTcpServer() {
	for _, s := range tcpList {
		s.Close()
		zap.L().Sugar().Infof(" [INFO] TcpProxyStop:%s\n", s.Addr)
	}
}
