package azure

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CreateContainer(containerName string) {

}

func DeleteContainer(containerName string) {

}

func GetContainers(endpoint string, token string) (containers []string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	return containers
}
