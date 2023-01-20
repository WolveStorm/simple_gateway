package grpc_middlewares

import (
	"errors"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func GRPCAPPFlowLimit(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		arr := md.Get("userid")
		arr1 := md.Get("qps")
		appId := ""
		qps := ""
		if len(arr) > 0 {
			appId = arr[0]
		} else {
			return errors.New("请先登录验证通过")
		}
		if len(arr1) > 0 {
			qps = arr1[0]
		}
		Qps, _ := strconv.Atoi(qps)
		// 对用户限流,限定用户名称
		userQPS := Qps
		if userQPS > 0 {
			clientLimiter := cache.FlowLimitManager.NewLimiter(global.ServiceLimiterKey+appId, int(userQPS))
			if !clientLimiter.Limit.Allow() {
				return errors.New("您已经被服务器限流")
			}
		}
		err := handler(srv, ss)
		if err != nil {
			return err
		}
		return nil
	}
}
