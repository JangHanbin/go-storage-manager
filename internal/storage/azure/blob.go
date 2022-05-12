package azure

import (
	"io"
)

func UploadBlob(containerName string, blobName string, data []byte) {

}

func PartialUploadBlob(containerName string, blobName string, reader io.Reader) {

}

func DownloadBlob(containerName string, blobName string, downloadedData *[]byte) {

}

func DeleteBlob(containerName string, blobName string) {

}

func GetBlobs(containerName string) (blobs []string) {

	return blobs
}
