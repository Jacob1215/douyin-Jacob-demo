package db

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

