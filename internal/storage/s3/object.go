package s3

import (
	"context"
	"io"
)

type S3PresignGetObjectAPI interface {
}

func GetObjects(bucketName string) (objects []string) {

	return objects
}

func DownloadObject(bucketName string, objectName string, downloadedData *[]byte) {

}

func UploadObject(bucketName string, fileName string, data []byte) {

}

func PartialUploadObject(bucketName string, fileName string, reader io.Reader) {

}

func DeleteObject(bucketName string, objectName string) {

}

func getPresignedURL(c context.Context, api S3PresignGetObjectAPI) {

}

func GetPublicURL(bucketName string, key string) (publicURL string) {

	return publicURL
}
