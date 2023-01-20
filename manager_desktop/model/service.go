package model

import "simple_gateway/global"

type ServiceDetail struct {
	Id            int            `json:"id"`
	ServiceInfo   *ServiceInfo   `json:"Info"`
	HttpRule      *HttpRule      `json:"Http"`
	TcpRule       *TcpRule       `json:"Tcp"`
	GrpcRule      *GrpcRule      `json:"GrpcRule"`
	LoadBalance   *LoadBalance   `json:"LoadBalance"`
	AccessControl *AccessControl `json:"AccessControl"`
}

type ServiceList struct {
	Total int                `json:"total"`
	List  []*ServiceListItem `json:"list"`
}

type ServiceListItem struct {
	Id          int    `json:"id"`
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
	LoadType    int    `json:"load_type"`
	ServiceAddr string `json:"service_addr"`
	QPS         int    `json:"qps"`
	QPD         int    `json:"qpd"` // 日请求量
	TotalNode   int    `json:"total_node"`
}

func GetServiceDetail(info *ServiceInfo) (*ServiceDetail, error) {
	lb := &LoadBalance{}
	err := lb.GetLoadBalanceById(int(info.ID))
	if err != nil {
		return nil, err
	}
	ac := &AccessControl{}
	err = ac.GetAccessControlById(int(info.ID))
	if err != nil {
		return nil, err
	}
	detail := &ServiceDetail{Id: int(info.ID)}
	detail.ServiceInfo = info
	detail.LoadBalance = lb
	detail.AccessControl = ac
	if info.LoadType == global.HTTPMode {
		httpRule := &HttpRule{}
		err := httpRule.GetHTTPRuleById(int(info.ID))
		if err != nil {
			return nil, err
		}
		detail.HttpRule = httpRule
	} else if info.LoadType == global.TCPMode {
		tcpRule := &TcpRule{}
		err := tcpRule.GetTCPRuleById(int(info.ID))
		if err != nil {
			return nil, err
		}
		detail.TcpRule = tcpRule
	} else if info.LoadType == global.GRPcMode {
		grpcRule := &GrpcRule{}
		err := grpcRule.GetGRPCRuleById(int(info.ID))
		if err != nil {
			return nil, err
		}
		detail.GrpcRule = grpcRule
	}
	return detail, nil
}
