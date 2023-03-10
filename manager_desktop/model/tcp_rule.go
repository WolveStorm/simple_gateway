package model

import "simple_gateway/global"

type TcpRule struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	ServiceID int64 `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port      int   `json:"port" gorm:"column:port" description:"端口	"`
}

func (t *TcpRule) TableName() string {
	return "gateway_service_tcp_rule"
}

func (t *TcpRule) GetTCPRuleById(id int) error {
	err := global.GORMClient.Where("service_id = ?", id).First(t).Error
	if err != nil {
		return err
	}
	return nil
}
