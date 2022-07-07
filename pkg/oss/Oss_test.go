package oss

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	CreateBucket("jacob-files-douyin")
}

func TestUploadLocalFile(t *testing.T) {
	err:= UploadLocalFile("jacob-files","test.mp4","./test.mp4")
	fmt.Println(err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()
	//fi, _ := os.Stat("./test.mp4")
	err := UploadFile("jacob-files","test2",file)
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url,err := GetFileUrl("jacob-files","test.mp4")
	fmt.Println(url,err,strings.Split(string(url),"?")[0])
}