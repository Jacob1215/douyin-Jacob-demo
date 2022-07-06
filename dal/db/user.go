package db

type User struct {
	BaseModel
	UserName string `gorm:"index:idx_username;unique;type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
	FavoriteVideo []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowCount int64 `gorm:"type:int;default:0;not null" json:"follow_count"`
	FollowerCount int64 `gorm:"type:int;default:0;not null" json:"follower_count"`

	IsFollow bool `gorm:"default:false;not null" json:"is_follow"`
}