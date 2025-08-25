// @author xiangqian
// @date 2025/08/01 13:13
package minio

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	err := Init(Config{
		Endpoint:  "localhost:9000",
		AccessKey: "KxlBMiNQdqbAmDKJ7q1f",
		SecretKey: "bqu9AIjArubbXUYwVm5hLbQNGlrNwh9QirCFCoQ9",
	})
	if err != nil {
		panic(err)
	}
}

func TestListBuckets(t *testing.T) {
	TestInit(t)
	bucketInfos, err := ListBuckets()
	if err != nil {
		panic(err)
	}
	fmt.Println(bucketInfos)
}

func TestBucketExists(t *testing.T) {
	TestInit(t)
	exists, err := BucketExists("test1")
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)
}

func TestMakeBucket(t *testing.T) {
	TestInit(t)
	err := MakeBucket("test")
	if err != nil {
		panic(err)
	}
}

func TestListObjects(t *testing.T) {
	TestInit(t)
	objInfoCh := ListObjects("test", "", true)
	for objInfo := range objInfoCh {
		if objInfo.Err != nil {
			fmt.Println(objInfo.Err)
			continue
		}
		fmt.Println(objInfo.Key)
	}
	fmt.Println()

	objInfoCh = ListObjects("test", "tmp", true)
	for objInfo := range objInfoCh {
		if objInfo.Err != nil {
			fmt.Println(objInfo.Err)
			continue
		}
		fmt.Println(objInfo.Key)
	}
}

func TestStatObject(t *testing.T) {
	TestInit(t)
	objInfo, err := StatObject("test", "afe87714e1b53c5ef493819518d64a1b")
	if err != nil {
		panic(err)
	}
	fmt.Println(objInfo)
	fmt.Println()

	objInfo, err = StatObject("test", "tmp/minio.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(objInfo)
}

func TestGetObject(t *testing.T) {
	TestInit(t)
	obj, err := GetObject("test", "tmp/minio.png")
	if err != nil {
		panic(err)
	}
	defer obj.Close()

	// 读取对象内容
	data, err := io.ReadAll(obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d\n", len(data))
}

func TestPresignedHeadObject(t *testing.T) {
	TestInit(t)
	url, err := PresignedHeadObject("test", "tmp/minio.png", 30*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}

func TestPresignedGetObject(t *testing.T) {
	TestInit(t)
	url, err := PresignedGetObject("test", "tmp/minio.png", 30*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}

func TestPresignedPutObject(t *testing.T) {
	TestInit(t)
	url, err := PresignedPutObject("test", "tmp/test.png", 30*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}

func TestPutObject(t *testing.T) {
	TestInit(t)

	data, err := os.ReadFile("C:\\Users\\xiangqian\\Pictures\\kelin.jpg")
	if err != nil {
		panic(err)
	}

	uploadInfo, err := PutObject("test", "tmp/kelin.jpg",
		bytes.NewReader(data),
		int64(len(data)),
		"image/jpeg")
	if err != nil {
		panic(err)
	}
	fmt.Println(uploadInfo)
}

func TestRemoveObject(t *testing.T) {
	TestInit(t)
	err := RemoveObject("test", "tmp/kelin.jpg")
	if err != nil {
		panic(err)
	}
}
func TestRemoveBucket(t *testing.T) {
	TestInit(t)
	err := RemoveBucket("test")
	if err != nil {
		panic(err)
	}
}
