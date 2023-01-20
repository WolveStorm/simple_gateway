package grpc_middlewares

import (
	"errors"
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"sync/atomic"
)

func GRPCAPPFlowCount(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		arr := md.Get("userid")
		arr1 := md.Get("qpd")
		appId := ""
		qpd := ""
		if len(arr) > 0 {
			appId = arr[0]
		} else {
			return errors.New("请先登录验证通过")
		}
		if len(arr1) > 0 {
			qpd = arr1[0]
		}
		Qpd, _ := strconv.Atoi(qpd)
		counter := cache.FlowManager.GetFlowCounter(global.UserFlowLimit + appId)
		atomic.AddInt64(&counter.AddTicket, 1)
		// 用户日请求量限制
		if Qpd > 0 {
			if int64(Qpd) <= counter.LastTicket {
				return errors.New("您的日请求量已经达到了上限")
			}
		}
		err := handler(srv, ss)
		if err != nil {
			return err
		}
		return nil
	}
}
