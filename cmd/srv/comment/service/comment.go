package service

import (
	"context"
	"douyin-Jacob/cmd/comment/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Comment) DouyinCommentAction(ctx context.Context,req *proto.DouyinCommentActionRequest) (*proto.DouyinCommentActionResponse,error) {
	//TODO
	comment := &db.Comment{
		UserID: req.UserId,
		VideoID: req.VideoId,
		Content: req.CommentText,
	}
	if req.ActionType == 1{
		//执行创建评论的逻辑
		err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			//新增评论数据
			err := tx.Create(comment).Error
			if err != nil{
				return err
			}
			//改变com count
			res :=tx.Model(&db.Video{}).Where("ID = ? ",comment.VideoID).Update("com_count",gorm.Expr("com_count + ? ",1))
			if err.Error != nil{
				return res.Error
			}
			if res.RowsAffected  != 1{
				return status.Errorf(codes.DataLoss,"cannot update video comment count")
			}
			return nil
		})
		if err != nil{
			return nil, err
		}
	}
	//进行删除评论操作
	if req.ActionType == 2{
		err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err := tx.First(&comment,req.CommentId).Error;err != nil{
				return err
			}
			//删除评论
			err := tx.Unscoped().Delete(&comment).Error
			if err != nil{
				return err
			}
			//改变video中com count
			res := tx.Model(&db.Video{}).Where("ID = ?",comment.VideoID).Update("com_count",gorm.Expr("com_count - ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1 {
				return status.Errorf(codes.DataLoss,"cannot update video comment count")
			}
			return nil
		})
		if err !=nil{
			return nil,err
		}
	}
	return &proto.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg: "comment action successed",
		Comment: &proto.Comment{
			Id: req.CommentId,
			User: &proto.User{
				Id: req.UserId,
			},
			//CreateDate: req.,//TODO
			Content: req.CommentText,
		},
	},nil
}



