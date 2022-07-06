package db

type Video struct {
	BaseModel
	Author User `gorm:"foreignkey:AuthorID"`
	AuthorID int64 `gorm:"index:idx_authorid;not null"`

 	PlayUrl string `gorm:"type:varchar(200);not null"`
	CoverUrl string `gorm:"type:varchar(200);not null"`

	FavCount int64 `gorm:"type:int;default:0;not null"`
	ComCount int64 `gorm:"type:int;default:0;not null"`

	IsFavorite bool `gorm:"type:bool;default:false;not null"`

	Data  []byte `gorm:"column:video_data"`
	Title string `gorm:"type:varchar(50);not null"`
}
