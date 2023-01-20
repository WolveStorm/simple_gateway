package service

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"simple_gateway/dto"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"simple_gateway/model"
	"simple_gateway/protoc"
	"simple_gateway/util"
	"strings"
)

func AddHTTPService(req form.AddHTTPServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.HTTPMode,
			ServiceDesc: req.ServiceDesc,
		}
		// 判断是否有重复名
		t := global.GORMClient.Where("service_name = ?", serviceInfo.ServiceName).Where("is_delete = ?", 0).First(&model.ServiceInfo{})
		if t.Error == nil {
			return errors.New("已经存在相同服务名了")
		}
		if t.Error != nil && t.Error != gorm.ErrRecordNotFound {
			return t.Error
		}
		err := global.GORMClient.Create(serviceInfo).Error
		if err != nil {
			return err
		}
		var serviceId = serviceInfo.ID
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			ForbidList:             req.ForbidList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Create(lbInfo).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Create(acInfo).Error
		if err != nil {
			return err
		}
		rule := &model.HttpRule{
			ServiceID:      int64(serviceId),
			RuleType:       req.RuleType,
			Rule:           req.Rule,
			NeedHttps:      req.NeedHttps,
			NeedWebsocket:  req.NeedWebsocket,
			NeedStripUri:   req.NeedStripUrl,
			UrlRewrite:     req.UrlRewrite,
			HeaderTransfor: req.HeaderTransfor,
		}
		// 判断有没有重复的HTTP规则配置
		t = global.GORMClient.Where("rule = ?", req.Rule).Where("rule_type = ?", req.RuleType).Find(&model.HttpRule{})
		if tx.RowsAffected != 0 {
			return errors.New("已经存在相同http配置")
		}
		if tx.Error != nil {
			return tx.Error
		}
		err = global.GORMClient.Create(rule).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			HttpRule:      rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateHTTPService(req form.UpdateHTTPServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.HTTPMode,
			ServiceDesc: req.ServiceDesc,
		}
		err := global.GORMClient.Table(serviceInfo.TableName()).Where("id = ?", req.Id).Updates(util.StructToUpdateMap(serviceInfo)).Error
		if err != nil {
			return err
		}
		var serviceId = req.Id
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			ForbidList:             req.ForbidList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Table(lbInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(lbInfo)).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Table(acInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(acInfo)).Error
		if err != nil {
			return err
		}
		rule := &model.HttpRule{
			ServiceID:      int64(serviceId),
			RuleType:       req.RuleType,
			Rule:           req.Rule,
			NeedHttps:      req.NeedHttps,
			NeedWebsocket:  req.NeedWebsocket,
			NeedStripUri:   req.NeedStripUrl,
			UrlRewrite:     req.UrlRewrite,
			HeaderTransfor: req.HeaderTransfor,
		}
		// Update attributes with `struct`, will only update non-zero fields
		err = global.GORMClient.Table(rule.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(rule)).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			HttpRule:      rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func AddTCPService(req form.AddTCPServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.TCPMode,
			ServiceDesc: req.ServiceDesc,
		}
		// 判断是否有重复名
		t := global.GORMClient.Where("service_name = ?", serviceInfo.ServiceName).Where("is_delete = ?", 0).First(&model.ServiceInfo{})
		if t.Error == nil {
			return errors.New("已经存在相同服务名了")
		}
		if t.Error != nil && t.Error != gorm.ErrRecordNotFound {
			return t.Error
		}
		err := global.GORMClient.Create(serviceInfo).Error
		if err != nil {
			return err
		}
		var serviceId = serviceInfo.ID
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Create(lbInfo).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Create(acInfo).Error
		if err != nil {
			return err
		}
		rule := &model.TcpRule{
			ServiceID: int64(serviceId),
			Port:      req.Port,
		}
		// 判断有没有重复的HTTP规则配置
		t = global.GORMClient.Where("port = ?", rule.Port).Find(&model.TcpRule{})
		if tx.RowsAffected != 0 {
			return errors.New("已经存在相同tcp配置")
		}
		if tx.Error != nil {
			return tx.Error
		}
		err = global.GORMClient.Create(rule).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			TcpRule:       rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateTCPService(req form.UpdateTCPServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.TCPMode,
			ServiceDesc: req.ServiceDesc,
		}
		err := global.GORMClient.Table(serviceInfo.TableName()).Where("id = ?", req.Id).Updates(util.StructToUpdateMap(serviceInfo)).Error
		if err != nil {
			return err
		}
		var serviceId = req.Id
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Table(lbInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(lbInfo)).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Table(acInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(acInfo)).Error
		if err != nil {
			return err
		}
		rule := &model.TcpRule{
			ServiceID: int64(serviceId),
			Port:      req.Port,
		}
		// Update attributes with `struct`, will only update non-zero fields
		err = global.GORMClient.Table(rule.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(rule)).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			TcpRule:       rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func AddGRPCService(req form.AddGRPCServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.GRPcMode,
			ServiceDesc: req.ServiceDesc,
		}
		// 判断是否有重复名
		t := global.GORMClient.Where("service_name = ?", serviceInfo.ServiceName).Where("is_delete = ?", 0).First(&model.ServiceInfo{})
		if t.Error == nil {
			return errors.New("已经存在相同服务名了")
		}
		if t.Error != nil && t.Error != gorm.ErrRecordNotFound {
			return t.Error
		}
		err := global.GORMClient.Create(serviceInfo).Error
		if err != nil {
			return err
		}
		var serviceId = serviceInfo.ID
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Create(lbInfo).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Create(acInfo).Error
		if err != nil {
			return err
		}
		rule := &model.GrpcRule{
			ServiceID:      int64(serviceId),
			Port:           req.Port,
			HeaderTransfor: req.HeaderTransfor,
		}
		// 判断有没有重复的HTTP规则配置
		t = global.GORMClient.Where("port = ?", rule.Port).Find(&model.GrpcRule{})
		if tx.RowsAffected != 0 {
			return errors.New("已经存在相同http配置")
		}
		if tx.Error != nil {
			return tx.Error
		}
		err = global.GORMClient.Create(rule).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			GrpcRule:      rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateGRPCService(req form.UpdateGRPCServiceReq) error {
	err := global.GORMClient.Transaction(func(tx *gorm.DB) error {
		serviceInfo := &model.ServiceInfo{
			ServiceName: req.ServiceName,
			LoadType:    global.GRPcMode,
			ServiceDesc: req.ServiceDesc,
		}
		err := global.GORMClient.Table(serviceInfo.TableName()).Where("id = ?", req.Id).Updates(util.StructToUpdateMap(serviceInfo)).Error
		if err != nil {
			return err
		}
		var serviceId = req.Id
		lbInfo := &model.LoadBalance{
			ServiceID:              int64(serviceId),
			CheckMethod:            req.CheckMethod,
			CheckTimeout:           req.CheckTimeout,
			CheckInterval:          req.CheckInterval,
			RoundType:              req.RoundType,
			IpList:                 req.IPList,
			WeightList:             req.WeightList,
			UpstreamConnectTimeout: req.UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  req.UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    req.UpstreamIdleTimeout,
			UpstreamMaxIdle:        req.UpstreamMaxIdle,
		}
		// 判断权重列表的个数是否等于IP列表的个数
		if len(strings.Split(req.IPList, ",")) != len(strings.Split(req.WeightList, ",")) {
			return errors.New("权重列表和IP列表个数不一致")
		}
		err = global.GORMClient.Table(lbInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(lbInfo)).Error
		if err != nil {
			return err
		}
		acInfo := &model.AccessControl{
			ServiceID:         int64(serviceId),
			OpenAuth:          req.OpenAuth,
			BlackList:         req.BlackList,
			WhiteList:         req.WhiteList,
			ClientIPFlowLimit: req.ClientFlowLimit,
			ServiceFlowLimit:  req.ServiceFlowLimit,
		}
		err = global.GORMClient.Table(acInfo.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(acInfo)).Error
		if err != nil {
			return err
		}
		rule := &model.GrpcRule{
			ServiceID:      int64(serviceId),
			Port:           req.Port,
			HeaderTransfor: req.HeaderTransfor,
		}
		// Update attributes with `struct`, will only update non-zero fields
		err = global.GORMClient.Table(rule.TableName()).Where("service_id = ?", req.Id).Updates(util.StructToUpdateMap(rule)).Error
		if err != nil {
			return err
		}
		detail := model.ServiceDetail{
			Id:            int(serviceId),
			ServiceInfo:   serviceInfo,
			GrpcRule:      rule,
			LoadBalance:   lbInfo,
			AccessControl: acInfo,
		}
		model.ServiceSync(detail)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(req form.DeleteHTTPServiceReq) error {
	info := &model.ServiceInfo{}
	tx := global.GORMClient.Find(info, "id = ?", req.Id)
	if tx.RowsAffected == 0 {
		return errors.New("不存在该服务")
	}
	if tx.Error != nil {
		return tx.Error
	}
	err := global.GORMClient.Table(info.TableName()).Where("id = ?", info.ID).Update("is_delete", 1).Error
	if err != nil {
		return err
	}
	model.DeleteServiceSync(info.ServiceName)
	return nil
}

func ServiceDetail(req form.ServiceDetailReq) (*model.ServiceDetail, error) {
	info := &model.ServiceInfo{}
	err := info.GetServiceInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	lb := &model.LoadBalance{}
	err = lb.GetLoadBalanceById(req.Id)
	if err != nil {
		return nil, err
	}
	ac := &model.AccessControl{}
	err = ac.GetAccessControlById(req.Id)
	if err != nil {
		return nil, err
	}
	detail := &model.ServiceDetail{Id: req.Id}
	detail.ServiceInfo = info
	detail.LoadBalance = lb
	detail.AccessControl = ac
	if info.LoadType == global.HTTPMode {
		httpRule := &model.HttpRule{}
		err := httpRule.GetHTTPRuleById(req.Id)
		if err != nil {
			return nil, err
		}
		detail.HttpRule = httpRule
	} else if info.LoadType == global.TCPMode {
		tcpRule := &model.TcpRule{}
		err := tcpRule.GetTCPRuleById(req.Id)
		if err != nil {
			return nil, err
		}
		detail.TcpRule = tcpRule
	} else if info.LoadType == global.GRPcMode {
		grpcRule := &model.GrpcRule{}
		err := grpcRule.GetGRPCRuleById(req.Id)
		if err != nil {
			return nil, err
		}
		detail.GrpcRule = grpcRule
	}
	return detail, nil
}

func ServiceList(req form.ServiceListReq) (*model.ServiceList, error) {
	info := &model.ServiceInfo{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	all, total := info.PageSelect(req)
	offset := (req.PageNum - 1) * req.PageSize
	limit := offset + req.PageSize
	if limit > total {
		limit = total
	}
	infoList := make([]*model.ServiceListItem, 0)
	list := &model.ServiceList{Total: total}
	for _, v := range all[offset:limit] {
		dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
		client := protoc.NewFlowCountClient(dial)
		rsp, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: v.ServiceName})
		if err != nil {
			return nil, err
		}
		item := &model.ServiceListItem{
			Id:          int(v.ID),
			ServiceName: v.ServiceName,
			ServiceDesc: v.ServiceDesc,
			LoadType:    v.LoadType,
			ServiceAddr: v.GetServiceAddr(),
			QPS:         int(rsp.Qps),
			QPD:         int(rsp.Qpd),
			TotalNode:   v.GetNodeCount(),
		}
		infoList = append(infoList, item)
	}
	list.List = infoList
	return list, nil
}

func ServiceStat(req form.ServiceStatReq) (*dto.ServiceStatOutput, error) {
	info := &model.ServiceInfo{}
	err := info.GetServiceInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
	client := protoc.NewFlowCountClient(dial)
	rsp, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: info.ServiceName})
	if err != nil {
		return nil, err
	}
	output := &dto.ServiceStatOutput{
		Info: dto.ServiceInfo{ID: info.ID},
	}

	output.Today = rsp.TodayCount
	output.Yesterday = rsp.YesterdayCount
	return output, nil
}

func DashboardFlowStat() (*dto.DashboardStatOutput, error) {
	dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
	client := protoc.NewFlowCountClient(dial)
	rsp, err := client.GetServiceFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: global.TotalKey})
	if err != nil {
		return nil, err
	}
	output := &dto.DashboardStatOutput{}

	output.Today = rsp.TodayCount
	output.Yesterday = rsp.YesterdayCount
	return output, nil
}
