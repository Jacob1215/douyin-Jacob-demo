package db

type Relation struct {
	BaseModel
	User User `gorm:"foreignkey:UserID;"`
	UserID int64 `gorm:"index:idx_userid,unique;not null"`
	ToUser User `gorm:"foreignkey:ToUserID;"`
	ToUserID int64 `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}