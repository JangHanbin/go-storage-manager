package storage

import (
	"github.com/janghanbin/go-storage-manager/configs"
	hubAWS "github.com/janghanbin/go-storage-manager/internal/storage/azure"
	hubAzure "github.com/janghanbin/go-storage-manager/internal/storage/s3"
	"io"
)

type Storage interface {
	CreateList(string)
	GetList() []string
	GetFileList(string) []string
	UploadFile(string, string, []byte)
	PartialUploadFile(string, string, io.Reader)
	DownloadFile(string, string, *[]byte)
}
type Azure struct {
	data []byte
}

type S3 struct {
	data []byte
}

func (a *Azure) CreateList(containerName string) {
	hubAzure.CreateContainer(containerName)
}

func (a *Azure) GetList() []string {
	return hubAzure.GetContainers()
}

func (a *Azure) GetFileList(containerName string) []string {
	return hubAzure.GetBlobs(containerName)
}

func (a *Azure) UploadFile(containerName string, blobName string, data []byte) {
	hubAzure.UploadBlob(containerName, blobName, data)
}

func (a *Azure) DownloadFile(containerName string, blobName string, data *[]byte) {
	hubAzure.DownloadBlob(containerName, blobName, &a.data)
	data = &a.data
}
func (a *Azure) PartialUploadFile(containerName string, blobName string, i io.Reader) {
	hubAzure.PartialUploadBlob(containerName, blobName, i)
}

func (s *S3) CreateList(bucketName string) {
	hubAWS.CreateBucket(bucketName, "ap-northeast-2")
}

func (s *S3) GetList() []string {
	return hubAWS.GetBuckets()
}

func (s *S3) GetFileList(bucketName string) []string {
	return hubAWS.GetObjects(bucketName)
}

func (s *S3) UploadFile(bucketName string, objectName string, data []byte) {
	hubAWS.UploadObject(bucketName, objectName, data)
}

func (s *S3) DownloadFile(bucketName string, objectName string, data *[]byte) {
	hubAWS.DownloadObject(bucketName, objectName, &s.data)
	data = &s.data
}

func (s *S3) PartialUploadFile(bucketName string, objectName string, i io.Reader) {
	hubAWS.PartialUploadObject(bucketName, objectName, i)
}

func NewClient(client string, cfg *configs.Configuration) Storage {
	if client == "azure" {
		return &Azure{
			hubAzure.GetClient(cfg.Azure),
		}
	}
	if client == "aws" {
		return &S3{
			hubAWS.GetClient(cfg.AWS),
		}
	}
	// default
	return &Azure{
		hubAzure.GetClient(cfg.Azure),
	}
}
