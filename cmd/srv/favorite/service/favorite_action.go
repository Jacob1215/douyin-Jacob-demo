package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/favorite/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)
//点赞和删除操作。//注意查表的操作
func (s *Favorite) DouyinFavoriteAction(ctx context.Context, req *proto.DouyinFavoriteActionRequest)(*proto.DouyinFavoriteActionResponse,error){
	if req.ActionType == 1{
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			user := new(db.User)
			//找到user
			if err := tx.WithContext(ctx).First(user,req.UserId).Error;err != nil{
				return err
			}
			//找到视频
			video := new(db.Video)
			if err := tx.WithContext(ctx).First(video,req.VideoId).Error;err !=nil{
				return err
			}
			//把用户加到视频后对面
			if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideo").Append(video);err != nil{
				return err
			}
			//改变video表中的favcount
			res := tx.Model(video).Update("fav_count",gorm.Expr("fav_count + ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1{
				return status.Errorf(codes.DataLoss, "database error")
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	//取消点赞
	if req.ActionType == 2{
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			user := new(db.User)
			if err := tx.WithContext(ctx).First(user,req.UserId).Error;err != nil{
				return err
			}
			//通过user找到video
			video := new(db.Video)
			if err := global2.DB.WithContext(ctx).Model(&user).Association("FavoriteVideo").Find(&video,req.VideoId);err !=nil{
				return err
			}
			//删除user表中的这个视频。
			err := tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideo").Delete(video)
			if err != nil{
				return err
			}
			//改变video表中的favconut
			res := tx.Model(video).Update("fav_count",gorm.Expr("fav_count - ? ",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1 {
				return status.Errorf(codes.DataLoss,"video delete fav count failed")
			}
			return nil
		})
		if err != nil{
			return nil,err
		}
	}
	return &proto.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg: "fav action successed",
	},nil
}

