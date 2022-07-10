package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/comment/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/pkg/errno"
	proto "douyin-Jacob/proto"
	"gorm.io/gorm"
)

func (s *Comment) DouyinCommentAction(ctx context.Context,req *proto.DouyinCommentActionRequest) (*proto.DouyinCommentActionResponse,error) {
	comment := &db.Comment{
		UserID: req.UserId,
		VideoID: req.VideoId,
		Content: req.CommentText,
	}
	if req.ActionType == 1{
		//执行创建评论的逻辑
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			//新增评论数据
			err := tx.Create(comment).Error
			if err != nil{
				return errno.ErrCreateModelErr
			}
			//改变com count
			res :=tx.Model(&db.Video{}).Where("ID = ? ",comment.VideoID).Update("com_count",gorm.Expr("com_count + ? ",1))
			if err.Error != nil{
				return errno.ErrUpdateModelErr
			}
			if res.RowsAffected  != 1{
				return errno.ErrRowsAffectedNotEquelToOne
			}
			return nil
		})
		if err != nil{
			return nil, err
		}
	}
	//进行删除评论操作
	if req.ActionType == 2{
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err := tx.First(&comment,req.CommentId).Error;err != nil{
				return err
			}
			//删除评论
			err := tx.Unscoped().Delete(&comment).Error
			if err != nil{
				return errno.ErrDeleteDate
			}
			//改变video中com count
			res := tx.Model(&db.Video{}).Where("ID = ?",comment.VideoID).Update("com_count",gorm.Expr("com_count - ?",1))
			if res.Error != nil{
				return errno.ErrUpdateModelErr
			}
			if res.RowsAffected != 1 {
				return errno.ErrRowsAffectedNotEquelToOne
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



