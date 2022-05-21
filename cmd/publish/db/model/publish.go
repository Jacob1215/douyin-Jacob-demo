package model

import (
"gorm.io/gorm"
"time"
)

type BaseModel struct {//自定义model，方便加上自己的字段。
	ID        int64          `gorm:"index:idx_id;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}

type VideoPublish struct {
	Video int64 `gorm:"primarykey;unique;type:int;not null"`
	User
	PlayUrl string `gorm:"type:varchar(200);not null"`
	CoverUrl string `gorm:"type:varchar(200);not null"`

	FavCount int64 `gorm:"type:int;not null"`
	ComCount int64 `gorm:"type:int;not null"`

	IsFavorite bool `gorm:"type:bool;default:false;not null"`

	Data  []byte `gorm:"column:video_data"`
	Title string `gorm:"type:varchar(200);not null"`
}


type User struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(100);not null"`

	FollowCount int64 `gorm:"type:int;not null"`
	FollowerCount int64 `gorm:"type:int;not null"`

	IsFollow bool `gorm:"default:false;not null"`
}