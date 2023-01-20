package model

import (
	"fmt"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"strings"
)

type ServiceInfo struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ServiceName string `json:"service_name"`
	LoadType    int    `json:"load_type"`
	ServiceDesc string `json:"service_desc"`
	BaseModel   `ignore:"true"`
}

func (ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (s *ServiceInfo) GetServiceInfoById(id int) error {
	err := global.GORMClient.Where("is_delete = ?", 0).First(s, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceInfo) SelectAll() ([]*ServiceInfo, int) {
	infos := make([]*ServiceInfo, 0)
	err := global.GORMClient.Where("is_delete = ?", 0).Find(&infos).Error
	if err != nil {
		return nil, 0
	}
	return infos, len(infos)
}

func (s *ServiceInfo) PageSelect(req form.ServiceListReq) ([]*ServiceInfo, int) {
	infos := make([]*ServiceInfo, 0)
	err := global.GORMClient.Where("is_delete = ?", 0).Where("service_name LIKE ?", "%"+req.Info+"%").Find(&infos).Error
	if err != nil {
		return nil, 0
	}
	return infos, len(infos)
}

func (s *ServiceInfo) GroupByMode() ([]*DashServiceStatItemOutput, error) {
	list := make([]*DashServiceStatItemOutput, 0)
	err := global.GORMClient.Table(s.TableName()).Select("load_type,count(*) as value").Where("is_delete = ?", 0).Group("load_type").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ServiceInfo) GetServiceAddr() string {
	var addr = ""
	cfg := global.DebugFullConfig.ClusterConfig
	if s.LoadType == global.HTTPMode {
		httpRule := &HttpRule{}
		err := httpRule.GetHTTPRuleById(int(s.ID))
		if err != nil {
			return ""
		}
		// 判断是http还是https
		if httpRule.NeedHttps == 1 {
			if httpRule.RuleType == global.PathType {
				addr = fmt.Sprintf("%s:%d%s", cfg.Host, cfg.SSLPort, httpRule.Rule)
			} else if httpRule.RuleType == global.DomainType {
				addr = fmt.Sprintf("%s", httpRule.Rule)
			}
		} else {
			if httpRule.RuleType == global.PathType {
				addr = fmt.Sprintf("%s:%d%s", cfg.Host, cfg.Port, httpRule.Rule)
			} else if httpRule.RuleType == global.DomainType {
				addr = fmt.Sprintf("%s", httpRule.Rule)
			}
		}
	} else if s.LoadType == global.TCPMode {
		tcpRule := &TcpRule{}
		err := tcpRule.GetTCPRuleById(int(s.ID))
		if err != nil {
			return ""
		}
		addr = fmt.Sprintf("%s:%d", cfg.Host, tcpRule.Port)
	} else if s.LoadType == global.GRPcMode {
		grpcRule := &GrpcRule{}
		err := grpcRule.GetGRPCRuleById(int(s.ID))
		if err != nil {
			return ""
		}
		addr = fmt.Sprintf("%s:%d", cfg.Host, grpcRule.Port)
	}
	return addr
}

func (s *ServiceInfo) GetNodeCount() int {
	lb := &LoadBalance{}
	err := lb.GetLoadBalanceById(int(s.ID))
	if err != nil {
		return 0
	}
	list := lb.IpList
	arr := strings.Split(list, ",")
	return len(arr)
}

type DashServiceStatItemOutput struct {
	Name     string `json:"name"`
	LoadType int    `json:"load_type"`
	Value    int64  `json:"value"`
}

type ServiceStat struct {
	Legend []string                     `json:"legend"`
	Series []*DashServiceStatItemOutput `json:"data"`
}
