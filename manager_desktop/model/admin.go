package model

import (
	"simple_gateway/global"
	"simple_gateway/util"
)

type AdminInfo struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	Salt     string
	Password string
	BaseModel
}

func (AdminInfo) TableName() string {
	return "gateway_admin"
}

func GetAdminInfo(username string) (*AdminInfo, error) {
	info := &AdminInfo{}
	err := global.GORMClient.Where("user_name = ?", username).Find(info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func ChangePassword(password string) error {
	info := &AdminInfo{}
	err := global.GORMClient.Where("user_name = ?", "admin").Find(info).Error
	if err != nil {
		return err
	}

	changePwd := util.GetSaltPassword(info.Salt, password)
	info.Password = changePwd

	err = global.GORMClient.Save(info).Error
	if err != nil {
		return err
	}
	return nil
}
