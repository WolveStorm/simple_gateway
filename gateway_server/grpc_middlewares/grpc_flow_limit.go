package grpc_middlewares

import (
	"errors"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"strings"
)

func GRPCFlowLimit(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// 对服务器限流，限定服务器
		serverQps := detail.AccessControl.ServiceFlowLimit
		if serverQps > 0 {
			serverLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName, serverQps)
			if !serverLimiter.Limit.Allow() {
				return errors.New("服务器访问次数达到上限")
			}
		}
		context, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("获得peer失败")
		}
		split := strings.Split(context.Addr.String(), ":")
		clientIp := ""
		if len(split) == 2 {
			clientIp = split[1]
		}
		// 对客户端限流,限定客户端IP
		clientQps := detail.AccessControl.ClientIPFlowLimit
		if clientQps > 0 {
			clientLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+detail.ServiceInfo.ServiceName+clientIp, clientQps)
			if !clientLimiter.Limit.Allow() {
				return errors.New("您今日内访问服务器的次数达到上限")
			}
		}
		err := handler(srv, ss)
		if err != nil {
			return err
		}
		return nil
	}
}
