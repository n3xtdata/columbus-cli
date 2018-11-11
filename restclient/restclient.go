package restclient

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
)

var BaseUrl = "http://localhost:8080/api"

func Get(path string) string {

	url := fmt.Sprintf(BaseUrl + "/" + path)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "nil"
	}
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Do: ", err)
		return "nil"
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
