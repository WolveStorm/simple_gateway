package model

import "simple_gateway/global"

type GrpcRule struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	ServiceID      int64  `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port           int    `json:"port" gorm:"column:port" description:"端口	"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (g *GrpcRule) TableName() string {
	return "gateway_service_grpc_rule"
}

func (g *GrpcRule) GetGRPCRuleById(id int) error {
	err := global.GORMClient.Where("service_id = ?", id).First(g).Error
	if err != nil {
		return err
	}
	return nil
}
