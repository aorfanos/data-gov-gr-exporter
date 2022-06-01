package sailing

import (
	"encoding/json"
	"log"

	"github.com/aorfanos/gov-gr-exporter/utils"
)

const ApiEndpointSailingTraffic string = "https://data.gov.gr/api/v1/query/sailing_traffic"

type SailingTraffic []struct {
	Date string `json:"date"`
	ArrivalPort string `json:"arrivalport"`
	ArrivalPortName string `json:"arrivalportname"`
	DeparturePort string `json:"departureport"`
	DeparturePortName string `json:"departureportname"`
	RouteCode string `json:"routecode"`
	RouteCodeName string `json:"routecodenames"`
	PassengerCount float64 `json:"passengercount"`
	VehicleCount float64 `json:"vehiclecount"`
}

func GetSailingTraffic() *SailingTraffic {
	dateFrom, dateTo := utils.GenerateDateToFrom()
	body, err := utils.NewGovGrGetRequest(ApiEndpointSailingTraffic+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var sailingTraffic SailingTraffic
	if err := json.Unmarshal(body, &sailingTraffic); err != nil {
		log.Fatal(err)
	}

	return &sailingTraffic
}

