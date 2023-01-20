package model

import (
	"errors"
	"gorm.io/gorm"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"simple_gateway/util"
	"time"
)

type User struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	AppID     string    `json:"app_id" gorm:"column:app_id" description:"租户id	"`
	Name      string    `json:"name" gorm:"column:name" description:"租户名称	"`
	Secret    string    `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string    `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配"`
	Qpd       int64     `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	Qps       int64     `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间" ignore:"true"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间" ignore:"true"`
	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是" ignore:"true"`
}

func (t *User) TableName() string {
	return "gateway_app"
}

func (u *User) PageSelect(req form.AppListReq) ([]*User, int) {
	infos := make([]*User, 0)
	err := global.GORMClient.Where("is_delete = ?", 0).Where("app_id LIKE ?", "%"+req.Info+"%").Find(&infos).Error
	if err != nil {
		return nil, 0
	}
	return infos, len(infos)
}

func (u *User) SelectAll() ([]*User, int) {
	infos := make([]*User, 0)
	err := global.GORMClient.Where("is_delete = ?", 0).Find(&infos).Error
	if err != nil {
		return nil, 0
	}
	return infos, len(infos)
}

func (u *User) Count() (int, error) {
	infos := make([]*User, 0)
	err := global.GORMClient.Find(&infos).Error
	if err != nil {
		return 0, err
	}
	return len(infos), nil
}

func (u *User) Save(req form.AddUserReq) error {
	u1 := &User{}
	err := global.GORMClient.Table(u.TableName()).Where("app_id = ? or name = ?", req.AppId, req.Name).Where("is_delete = ?", 0).First(u1).Error
	if err == nil {
		return errors.New("已经存在相同的appId")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	u.AppID = req.AppId
	if len(req.Secret) == 32 {
		u.Secret = req.Secret
	} else {

		u.Secret = util.MD5(req.AppId)
	}
	u.Qpd = int64(req.QPD)
	u.Qps = int64(req.QPS)
	u.Name = req.Name
	err = global.GORMClient.Save(u).Error
	if err != nil {
		return err
	}
	UserSync(*u)
	return nil
}

func (u *User) Update(req form.UpdateUserReq) error {
	u.AppID = req.AppId
	u.Secret = req.Secret
	u.Qpd = int64(req.QPD)
	u.Qps = int64(req.QPS)
	u.Name = req.Name
	err := global.GORMClient.Table(u.TableName()).Where("id = ?", req.Id).Updates(util.StructToUpdateMap(u)).Error
	if err != nil {
		return err
	}
	UserSync(*u)
	return nil
}

func (u *User) Delete(req form.DeleteUserReq) error {
	return global.GORMClient.Table(u.TableName()).Where("id = ?", req.Id).Update("is_delete", 1).Error
}

func (u *User) FindById(id int) error {
	err := global.GORMClient.Table(u.TableName()).First(u, id).Error
	if err != nil {
		return err
	}
	return nil
}

type UserList struct {
	Total int             `json:"total"`
	List  []*UserListItem `json:"list"`
}

type UserListItem struct {
	Id     int    `json:"id"`
	AppId  string `json:"app_id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	QPS    int    `json:"qps"`
	QPD    int    `json:"qpd"` // 日请求量
}
