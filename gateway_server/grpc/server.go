package grpc

import (
	"fmt"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"gateway_server/grpc_middlewares"
	"gateway_server/grpc_middlewares/proxy"
	"gateway_server/reverse_proxy/load_balance"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var grpcList []*grpc.Server

func StartGRPCServer() {
	service := cache.GetAllGRPCService()
	for _, v := range service {
		go func(detail *model.ServiceDetail) {
			lb := load_balance.GetLoadBalanceByService(detail)
			server := grpc.NewServer(
				grpc.ChainStreamInterceptor(
					grpc_middlewares.GRPCFlowCount(detail),
					grpc_middlewares.GRPCFlowLimit(detail),
					grpc_middlewares.GRPCJWTAuth(detail),
					grpc_middlewares.GRPCAPPFlowCount(detail),
					grpc_middlewares.GRPCAPPFlowLimit(detail),
					grpc_middlewares.GRPCWhiteList(detail),
					grpc_middlewares.GRPCBlackList(detail),
					grpc_middlewares.GrpcHeaderTransfor(detail),
				),
				grpc.CustomCodec(proxy.Codec()),
				grpc.UnknownServiceHandler(proxy.TransparentHandler(grpc_middlewares.GRPCReverseProxyWithLoadBalance(lb))))
			listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.ProxyFullConfig.GRPCConfig.Host, detail.GrpcRule.Port))
			if err != nil {
				zap.L().Sugar().Errorf("[Error] TcpProxyListenError:%d\n", detail.GrpcRule.Port)
			}
			err = server.Serve(listen)
			if err != nil {
				zap.S().Errorf("[Navi Gateway] start grpc proxy server error,port:%d", detail.GrpcRule.Port)
			}
			grpcList = append(grpcList, server)
		}(v)
	}
}

func CloseGrpcServer() {
	for _, server := range grpcList {
		server.Stop()
	}
}
