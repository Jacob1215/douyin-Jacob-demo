package service

import (
	global2 "douyin-Jacob/cmd/srv/user/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/pkg/oss_api/oss"
	"douyin-Jacob/proto"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/opentracing/opentracing-go"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm"
	"image"
	"image/jpeg"
	"os"

	"bytes"
	"context"
	"strings"
)

func (s *PublishServer) PostVideo(ctx context.Context,request *proto.DouyinPublishActionRequest)(*proto.DouyinPublishActionResponse,error) {
	//TODO 这里要做对象存储，还要做事务。
	ossVideoBucketName := oss.OssVideoBucketName


	reader := bytes.NewReader(request.Data)
	u2,err := uuid.NewV4() //给视频文件加编号
	if err !=nil{
		return  nil,err
	}
	fileName := u2.String()+"."+"mp4"
	//上传视频//这里想不明白。
	err = oss.UploadFile(ossVideoBucketName,fileName,reader)
	if err != nil{
		return nil, err
	}
	//获取视频连接
	urlDate,err := oss.GetFileUrl(ossVideoBucketName,fileName)
	playUrl := strings.Split(string(urlDate),"?")[0]
	if err != nil{
		return nil, err
	}
	u3,err := uuid.NewV4()
	if err !=nil{
		return nil, err
	}
	//获取封面
	coverPath := u3.String()+"."+"jpg"
	coverData,err := readFrameAsJpeg(playUrl)
	if err != nil{
		return nil, err
	}
	//上传封面
	coverReader := bytes.NewReader(coverData)
	err = oss.UploadFile(ossVideoBucketName,coverPath,coverReader)
	if err != nil{
		return nil, err
	}
	//获取封面连接
	coverUrl,err := oss.GetFileUrl(ossVideoBucketName,coverPath)
	if err != nil{
		return nil, err
	}
	cover := strings.Split(string(coverUrl),"?")[0]
	//封装


	videoModel := &db.Video{
		AuthorID: request.User.Id,
		PlayUrl: playUrl,
		CoverUrl: cover,
		FavCount: 0,
		ComCount: 0,
		Title: request.Title,
		Data: request.Data,
	}

	//事务+链路追踪
	parentSpan := opentracing.SpanFromContext(ctx)
	postVideoSpan := opentracing.GlobalTracer().StartSpan("post_video",opentracing.ChildOf(parentSpan.Context()))
	err = global2.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(videoModel).Error
		if err  != nil {
			return err
		}
		return nil
	})
	if err != nil{
		return nil,err
	}
	postVideoSpan.Finish()
	return &proto.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "publish video success",
	},nil

}

// ReadFrameAsJpeg//这个还不熟悉
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(url string) ([]byte,error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
