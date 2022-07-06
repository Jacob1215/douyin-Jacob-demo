package oss

import (
	global2 "douyin-Jacob/cmd/api/oss_api/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
)

var (
	OssClient *oss.Client
	endpoint = global2.ServerConfig.OssInfo.EndPoint
	OssAccessKeyId = global2.ServerConfig.OssInfo.ApiKey
	OssSecret = global2.ServerConfig.OssInfo.ApiSecrect
	OssVideoBucketName = global2.ServerConfig.OssInfo.UploadDir
)


func init()  {
	client,err := oss.New(endpoint, OssAccessKeyId, OssSecret)
	if err != nil{
		zap.S().Errorf("Oss client oss failed:%v",err)
	}
	err = CreateBucket(OssVideoBucketName)
	OssClient = client
	if err != nil{
		zap.S().Errorf("Oss create %v bucket failed:%v", OssVideoBucketName,err)
	}

}

func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0{
		zap.S().Error("Oss bucket name invalid")
	}
	// 创建名为examplebucket的存储空间，并设置存储类型为低频访问oss.StorageIA、读写权限ACL为公共读oss.ACLPublicRead、数据容灾类型为同城冗余存储oss.RedundancyZRS。
	err := OssClient.CreateBucket(bucketName,oss.StorageClass(oss.StorageIA),oss.ACL(oss.ACLPublicRead),oss.RedundancyType(oss.RedundancyZRS))
	if err != nil{
		exist,errBucketExists := OssClient.IsBucketExist(bucketName)
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
func UploadLocalFile(bucketName string,objectName string,filePath string)(error){
	//填写存储空间名称，
	bucket,err := OssClient.Bucket(bucketName)
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
func UploadFile(bucketName string,objectName string,reader io.Reader) error {
	//填写存储空间名称，
	bucket,err := OssClient.Bucket(bucketName)
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
func GetFileUrl(bucketName string,fileName string) ([]byte,error)  {

	//填写存储空间名称，
	bucket,err := OssClient.Bucket(bucketName)
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
