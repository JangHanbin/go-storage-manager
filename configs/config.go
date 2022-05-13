package configs

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	BLOBENDPOINT int = iota
	QUEUEENDPOINT
	FILEENDPOINT
	TABLEENDPOINT
	SIGNATURE
)

const (
	KEY int = iota
	VALUE
)

type Configuration struct {
	Azure AzureConfiguration
	AWS   AWSConfiguration
}

type AzureConfiguration struct {
	ConnectionString   string
	SASToken           string
	BlobServiceSASURL  string
	FileServiceSASURL  string
	QueueServiceSASURL string
	TableServiceSASURL string
}

type AWSConfiguration struct {
	AccessKey string
	SecretKey string
}

var (
	Cfg *Configuration
)

func ReadConfiguration(path string) *Configuration {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	Cfg = new(Configuration)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(Cfg)
	if err != nil {
		log.Fatal(err)
	}

	return Cfg
}

func GetAzureConnectionValue(connectionStrings string, connectionIdx int) (value string) {
	connectionString := strings.Split(connectionStrings, ";")[connectionIdx] // Azure connection string seperated by ';'
	keyValuePair := strings.Split(connectionString, "=")                     // Azure key value seperated by '='

	value = keyValuePair[VALUE]

	return value
}
