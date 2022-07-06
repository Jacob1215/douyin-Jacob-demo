package service

import (
	"context"
	"douyin-Jacob/cmd/favorite/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
)


func (s *Favorite)DouyinFavoriteList(ctx context.Context,req *proto.DouyinFavoriteListRequest)(*proto.DouyinFavoriteListResponse,error)  {
	favList := proto.DouyinFavoriteListResponse{}

	user := new(db.User)
	if err := global.DB.WithContext(ctx).First(user,req.UserId).Error;err != nil{
		return nil,err
	}
	videos := []db.Video{}
	if err := global.DB.WithContext(ctx).Model(&user).Association("FavoriteVideo").Find(&videos);err != nil{
		return nil,err
	}
	for _, video := range videos{
		favList.VideoList = append(favList.VideoList,&proto.Video{
			Author: &proto.User{
				Id: video.AuthorID,
				Name: video.Author.UserName,
				FollowCount: video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,

				IsFollow: video.Author.IsFollow,
			},
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavCount,
			CommentCount: video.ComCount,
			IsFavorite: video.IsFavorite,
			Title: video.Title,
		})
	}
	return &proto.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg: "get user fav list successed",
		VideoList: favList.VideoList,
	},nil
}
