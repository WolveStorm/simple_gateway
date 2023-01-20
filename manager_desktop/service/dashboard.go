package service

import (
	"context"
	"google.golang.org/grpc"
	"simple_gateway/dto"
	"simple_gateway/global"
	"simple_gateway/model"
	"simple_gateway/protoc"
)

func PanelGroupData() (*dto.PanelGroupData, error) {
	data := &dto.PanelGroupData{}
	info := model.ServiceInfo{}
	_, count := info.SelectAll()
	data.ServiceNum = count
	user := model.User{}
	appNum, err := user.Count()
	if err != nil {
		return nil, err
	}
	data.AppNum = appNum
	dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
	client := protoc.NewFlowCountClient(dial)
	rsp, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: global.TotalKey})
	if err != nil {
		return nil, err
	}
	data.TodayRequestNum = int(rsp.Qpd)
	data.CurrentQps = int(rsp.Qps)
	return data, nil
}

func DashboardServiceStat() (*model.ServiceStat, error) {
	stat := &model.ServiceStat{}
	legend := []string{}
	info := model.ServiceInfo{}
	groups, err := info.GroupByMode()
	if err != nil {
		return nil, err
	}
	nameMap := map[int]string{
		0: "HTTP",
		1: "TCP",
		2: "GRPC",
	}
	for i, v := range groups {
		groups[i].Name = nameMap[v.LoadType]
		legend = append(legend, groups[i].Name)
	}
	stat.Series = groups
	stat.Legend = legend
	return stat, nil
}
