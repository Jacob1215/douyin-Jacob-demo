package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {//自定义model，方便加上自己的字段。
	ID int64 `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Name string `gorm:"index:idx_name;unique;type:varchar(20);not null"`
	Password string `gorm:"type:varchar(100);not null"`

	FollowCount int64 `gorm:"type:int;not null"`
	FollowerCount int64 `gorm:"type:int;not null"`

	IsFollow bool `gorm:"default:false;not null"`
}
