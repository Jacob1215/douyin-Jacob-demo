package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
)

func CreateBucket(ossclient *oss.Client,bucketName string) error {
	if len(bucketName) <= 0{
		zap.S().Error("Oss bucket name invalid")
	}
	// 创建名为examplebucket的存储空间，并设置存储类型为低频访问oss.StorageIA、读写权限ACL为公共读oss.ACLPublicRead、数据容灾类型为同城冗余存储oss.RedundancyZRS。
	err := ossclient.CreateBucket(bucketName,oss.StorageClass(oss.StorageIA),oss.ACL(oss.ACLPublicRead),oss.RedundancyType(oss.RedundancyZRS))
	if err != nil{
		exist,errBucketExists := ossclient.IsBucketExist(bucketName)
		if errBucketExists  == nil && exist{
			zap.S().Errorf("bucket %v already exsits",bucketName)
			return nil
		} else {
			return err
		}
	} else {
		zap.S().Info("bucket create successfully")
	}
	return nil
}

// UploadLocalFile 上传本地文件（提供文件路径）至 Oss
func UploadLocalFile(ossclient *oss.Client,bucketName string,objectName string,filePath string)(error){
	//填写存储空间名称，
	bucket,err := ossclient.Bucket(bucketName)
	if err !=nil{
		zap.S().Errorf("get bucket failed")
		return err
	}
	//objectName是文件名，filePath是路径
	errPush := bucket.PutObjectFromFile(objectName,filePath)
	if errPush != nil{
		zap.S().Errorf("upload local file failed:%s",err)
		return err
	}
	zap.S().Infof("upload %s successfully ",objectName)
	return nil
}
//上传文件
func UploadFile(ossclient *oss.Client,bucketName string,objectName string,reader io.Reader) error {
	//填写存储空间名称，
	bucket,err := ossclient.Bucket(bucketName)
	if err !=nil{
		zap.S().Errorf("get bucket failed")
		return err
	}
	//将文件上传至文件中
	err = bucket.PutObject(objectName,reader)
	if err != nil{
		zap.S().Errorf("upload file failed:%v",err)
		return err
	}
	zap.S().Info("upload file succesed")
	return nil
}

//从OSS获取文件Url//用流式传输的方法?
func GetFileUrl(ossclient *oss.Client,bucketName string,fileName string) ([]byte,error)  {
	//填写存储空间名称，
	bucket,err := ossclient.Bucket(bucketName)
	if err !=nil{
		zap.S().Errorf("get bucket failed")
		return nil,err
	}
	//下载文件到流
	body,err := bucket.GetObject(fileName)
	if err != nil{
		zap.S().Errorf("get obj file %s failed:%v",fileName,err)
		return nil, err
	}
	defer body.Close()
	data,err := ioutil.ReadAll(body)
	if err !=nil{
		zap.S().Errorf("read obj file %s failed:%v",fileName,err)
		return nil,err
	}
	return data,nil
}
