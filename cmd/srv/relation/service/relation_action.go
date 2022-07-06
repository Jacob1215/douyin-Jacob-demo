package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/relation/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Relation) DouyinRelationAction(ctx context.Context,req *proto.DouyinRelationActionRequest)(*proto.DouyinRelationActionResponse,error) {
	//关注
	if req.ActionType == 1{
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			//新增关注数据
			err := tx.Create(&db.Relation{UserID: req.UserId,ToUserID: req.ToUserId}).Error
			if err != nil{
				return err
			}
			//改变user表中的following count
			res := tx.Model(new(db.User)).Where("ID = ?",req.UserId).Update("follow_count",gorm.Expr("follow_count + ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1{
				return status.Errorf(codes.DataLoss,"cannot update user follow relation")
			}
			//改变user表中的followercount？
			res = tx.Model(new(db.User)).Where("ID = ?",req.ToUserId).Update("follower_count",gorm.Expr("follower_count + ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1{
				return status.Errorf(codes.DataLoss,"cannot update user follower relation")
			}
			return nil
		})
		if err != nil{
			return nil, err
		}
	}
	//删除关注
	if req.ActionType == 2{
		err := global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			//删除关注数据
			relation := new(db.Relation)
			if err := tx.Where("user_id = ? AND to_user_id=?",req.UserId,req.ToUserId).First(&relation).Error;err != nil{
				return err
			}
			err := tx.Unscoped().Delete(&relation).Error
			if err != nil{
				return err
			}
			//改变user表中的following count
			res := tx.Model(new(db.User)).Where("ID = ?",req.UserId).Update("follow_count",gorm.Expr("follow_count - ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1{
				return status.Errorf(codes.DataLoss,"delete user follow failed")
			}
			//改变被关注者的数据
			res = tx.Model(new(db.User)).Where("ID = ?",req.ToUserId).Update("follower_count",gorm.Expr("follower_count - ?",1))
			if res.Error != nil{
				return res.Error
			}
			if res.RowsAffected != 1{
				return status.Errorf(codes.DataLoss,"delete user follower failed")
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return &proto.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "relation Action Successed",
	},nil
}