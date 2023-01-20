package model

import "simple_gateway/global"

// AccessControl
// 该配置用于配置用户使用网关限制的相关参数。
// 例如作用于哪个serviceId,是否进行token检查
// 白名单，黑名单
// server和client限流。
type AccessControl struct {
	ID                int64  `json:"id" gorm:"primary_key"`
	ServiceID         int64  `json:"service_id" gorm:"column:service_id" description:"服务id"`
	OpenAuth          int    `json:"open_auth" gorm:"column:open_auth" description:"是否开启权限 1=开启"`
	BlackList         string `json:"black_list" gorm:"column:black_list" description:"黑名单ip	"`
	WhiteList         string `json:"white_list" gorm:"column:white_list" description:"白名单ip	"`
	WhiteHostName     string `json:"white_host_name" gorm:"column:white_host_name" description:"白名单主机	"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流	"`
	ServiceFlowLimit  int    `json:"service_flow_limit" gorm:"column:service_flow_limit" description:"服务端限流	"`
}

func (a *AccessControl) TableName() string {
	return "gateway_service_access_control"
}

func (a *AccessControl) GetAccessControlById(id int) error {
	err := global.GORMClient.Where("service_id = ?", id).First(a).Error
	if err != nil {
		return err
	}
	return nil
}
