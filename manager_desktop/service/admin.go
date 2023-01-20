package service

import (
	"errors"
	"simple_gateway/model"
	"simple_gateway/util"
)

func Login(username, password string) (*model.AdminInfo, error) {
	info, err := model.GetAdminInfo(username)
	if err != nil {
		return nil, err
	}
	saltPassword := util.GetSaltPassword(info.Salt, password)
	if saltPassword != info.Password {
		return nil, errors.New("密码错误")
	}
	return info, nil
}

func ChangePassword(password string) error {
	return model.ChangePassword(password)
}
