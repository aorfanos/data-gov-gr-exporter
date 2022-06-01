package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func NewGovGrGetRequest(endpoint string) ([]byte, error) {
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Authorization", fmt.Sprintf("Token %s", os.Getenv("GOV_GR_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("Status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	jsonBody, err := ioutil.ReadAll(response.Body)
	return jsonBody, err
}