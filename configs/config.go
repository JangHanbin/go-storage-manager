package configs

import (
	"encoding/json"
	"log"
	"os"
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
