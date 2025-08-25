// @author xiangqian
// @date 2025/07/26 21:56
package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"time"
)

var client *minio.Client

func Init(config Config) error {
	var err error
	client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}

	_, err = client.ListBuckets(context.Background())
	return err
}

// ListBuckets 获取全部存储桶
func ListBuckets() ([]minio.BucketInfo, error) {
	return client.ListBuckets(context.Background())
}

// BucketExists 查看存储桶是否存在
// bucketName 存储桶名称
func BucketExists(bucketName string) (bool, error) {
	return client.BucketExists(context.Background(), bucketName)
}

// MakeBucket 创建存储桶
// bucketName 存储桶名称
func MakeBucket(bucketName string) error {
	return client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
		Region: "",
	})
}

// ListObjects 获取存储桶下的所有对象
// bucketName 存储桶名称
// bucketName 对象前缀
// recursive  是否递归查询
func ListObjects(bucketName, prefix string, recursive bool) <-chan minio.ObjectInfo {
	return client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: recursive,
	})
}

// StatObject 获取对象详情
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
func StatObject(bucketName, objectName string) (minio.ObjectInfo, error) {
	return client.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
}

// GetObject 获取对象流
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
func GetObject(bucketName, objectName string) (*minio.Object, error) {
	return client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
}

// PresignedHeadObject 生成预签名元数据对象地址（获取对象元数据而不需要下载对象内容）
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
// expire     失效时间
func PresignedHeadObject(bucketName, objectName string, expire time.Duration) (*url.URL, error) {
	return client.PresignedHeadObject(context.Background(), bucketName, objectName, expire, nil)
}

// PresignedGetObject 生成预签名下载对象地址
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
// expire     失效时间
func PresignedGetObject(bucketName, objectName string, expire time.Duration) (*url.URL, error) {
	return client.PresignedGetObject(context.Background(), bucketName, objectName, expire, nil)
}

// PresignedPutObject 生成预签名上传对象地址
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
// expire     失效时间
func PresignedPutObject(bucketName, objectName string, expire time.Duration) (*url.URL, error) {
	return client.PresignedPutObject(context.Background(), bucketName, objectName, expire)
}

// PutObject 上传对象
// bucketName  存储桶名称
// objectName  对象名称（包含路径）
// reader      数据源
// size        数据大小
// contentType 内容类型，例如：text/plain
func PutObject(bucketName, objectName string, reader io.Reader, size int64, contentType string) (minio.UploadInfo, error) {
	return client.PutObject(context.Background(), bucketName, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
}

// RemoveObject 删除对象
// bucketName 存储桶名称
// objectName 对象名称（包含路径）
func RemoveObject(bucketName, objectName string) error {
	return client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
}

// RemoveBucket 删除存储桶
// bucketName 存储桶名称
func RemoveBucket(bucketName string) error {
	return client.RemoveBucket(context.Background(), bucketName)
}

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
}
