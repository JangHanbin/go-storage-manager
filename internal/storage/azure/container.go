package azure

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func CreateContainer(containerName string) {

}

func DeleteContainer(containerName string) {

}

func GetContainers(endpoint string, token string) (containers []string) {
	u, _ := url.Parse(endpoint)

	u.RawQuery = token
	q := u.Query()
	q.Add("comp", "list")
	q.Add("restype", "container")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result EnumerationResults
	err = xml.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	for _, container := range result.Containers {
		containers = append(containers, container.Name) // append names into return value
	}

	return containers
}
