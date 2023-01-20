package model

import "time"

type BaseModel struct {
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete"`
}
