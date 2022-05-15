package configs

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
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
	SAS string
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

func SplitSAS(tokenString string) (endpoint string, token string) {
	u, err := url.Parse(tokenString)

	if err != nil {
		log.Fatal(err)
	}
	endpoint = u.Scheme + "://" + u.Opaque + u.Host + u.Path
	token = u.RawQuery
	return
}
