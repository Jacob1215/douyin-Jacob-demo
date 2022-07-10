package db

type Comment struct {
	BaseModel //这个basemodel跟其他的类似
	Video Video `gorm:"foreignkey:VideoID"`
	VideoID int64 `gorm:"index:idx_videoid;not null"`
	User User `gorm:"foreignkey:UserID"`
	UserID int64 `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}