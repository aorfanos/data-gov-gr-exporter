package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var dataGovGrRequestsCount = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "govgr_exporter_api_requests_total",
		Help: "The total number of requests to the data.gov.gr API.",
	},
)

func NewGovGrGetRequest(endpoint string) ([]byte, error) {
	dataGovGrRequestsCount.Inc()

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