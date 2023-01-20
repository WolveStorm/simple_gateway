package grpc_middlewares

import (
	"gateway_server/cache"
	"gateway_server/cache/model"
	"gateway_server/global"
	"google.golang.org/grpc"
	"sync/atomic"
)

func GRPCFlowCount(detail *model.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		counter := cache.FlowManager.GetFlowCounter(detail.ServiceInfo.ServiceName)
		atomic.AddInt64(&counter.AddTicket, 1)
		totalCounter := cache.FlowManager.GetFlowCounter(global.TotalKey)
		atomic.AddInt64(&totalCounter.AddTicket, 1)
		err := handler(srv, ss)
		if err != nil {
			return err
		}
		return nil
	}
}
