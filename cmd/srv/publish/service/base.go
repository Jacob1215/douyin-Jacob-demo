package service

import (
	"context"
	"douyin-Jacob/cmd/srv/publish/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/proto"
)

type PublishServer struct {
}

func Video(ctx context.Context,video *db.Video,fromID int64)(*proto.Video,error)  {
	if video == nil{
		return nil,errno.ErrVideoNotFound
	}
	//查询User是谁？
	user := new(db.User)
	if err := global.DB.WithContext(ctx).First(&user,video.AuthorID).Error ;err !=nil{
		return nil,err
	}
	//打包User信息
	if user == nil{
		return nil,errno.ErrUserNotFound
	}
	//查询用户favorite关系
	isFollow :=false
	relation := new(db.Relation)
	if err := global.DB.WithContext(ctx).First(&relation,"user_id = ? and to_user_id = ?",fromID,user.ID).Error; err != nil{
		return nil,errno.ErrFavRelationFailed
	}
	if relation != nil{
		isFollow = true
	}
	userInfo := &proto.User{
		Id: user.ID,
		Name: user.UserName,
		FollowCount: user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow: isFollow,
	}
	//打包视频信息
	return &proto.Video{
		Id: video.ID,
		Author: userInfo,
		PlayUrl: video.PlayUrl,
		CoverUrl: video.CoverUrl,
		FavoriteCount: video.FavCount,
		CommentCount: video.ComCount,
		Title: video.Title,
	},nil
}

//批量打包视频信息
func Videos(ctx context.Context,videos []*db.Video,fromID *int64)([]*proto.Video,error)  {
	vs := make([]*proto.Video,0)
	for _,v := range videos{
		video ,err := Video(ctx,v,*fromID)
		if err != nil{
			return nil, err
		}
		if video != nil {
			flag := false
			//
			if *fromID != 0 {
				userInfo := new(db.User)
				if err = global.DB.WithContext(ctx).First(userInfo, fromID).Error; err != nil {
					return nil, errno.ErrUserNotFound
				}
				videoInfo := new(db.Video)
				if err = global.DB.WithContext(ctx).Model(&userInfo).Association("FavoriteVideos").Find(&videoInfo, video.Id); err != nil {
					return nil, errno.ErrVideoNotFound
				}
				if videoInfo != nil && videoInfo.AuthorID != 0 {
					flag = true
				}
			}
			video.IsFavorite = flag
			vs = append(vs, video)
		}
	}
	return vs,nil
}