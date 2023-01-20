package grpc_impl

import (
	"context"
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/protoc"
	"time"
)

type FlowCountService struct{}

func (*FlowCountService) GetServiceFlowCount(ctx context.Context, req *protoc.FlowCountRequest) (*protoc.FlowCountResponse, error) {
	serviceName := req.ServiceName
	counter := cache.FlowManager.GetFlowCounter(serviceName)
	rsp := &protoc.FlowCountResponse{
		Qpd: cache.GetDayData(serviceName, time.Now()),
		Qps: int32(counter.QPS),
	}
	location, _ := time.LoadLocation("Asia/Chongqing")
	today := make([]int32, 0)
	nowTime := time.Now()
	h := nowTime.Hour()
	for i := 0; i <= h; i++ {
		today = append(today, cache.GetHourData(serviceName, time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), i, 0, 0, 0, location)))
	}
	yesterday := make([]int32, 0)
	lastTime := time.Now().Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i < 24; i++ {
		yesterday = append(yesterday, cache.GetHourData(serviceName, time.Date(lastTime.Year(), lastTime.Month(), lastTime.Day(), i, 0, 0, 0, location)))
	}
	rsp.TodayCount = today
	rsp.YesterdayCount = yesterday
	return rsp, nil
}
func (*FlowCountService) GetUserFlowCount(ctx context.Context, req *protoc.FlowCountRequest) (*protoc.FlowCountResponse, error) {
	userName := req.ServiceName
	counter := cache.FlowManager.GetFlowCounter(global.UserFlowLimit + userName)
	rsp := &protoc.FlowCountResponse{
		Qpd: int32(cache.GetDayData(global.UserFlowLimit+userName, time.Now())),
		Qps: int32(counter.QPS),
	}
	location, _ := time.LoadLocation("Asia/Chongqing")
	today := make([]int32, 0)
	nowTime := time.Now()
	h := nowTime.Hour()
	for i := 0; i <= h; i++ {
		today = append(today, cache.GetHourData(global.UserFlowLimit+userName, time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), i, 0, 0, 0, location)))
	}
	yesterday := make([]int32, 0)
	lastTime := time.Now().Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i < 24; i++ {
		yesterday = append(yesterday, cache.GetHourData(global.UserFlowLimit+userName, time.Date(lastTime.Year(), lastTime.Month(), lastTime.Day(), i, 0, 0, 0, location)))
	}
	rsp.TodayCount = today
	rsp.YesterdayCount = yesterday
	return rsp, nil
}
