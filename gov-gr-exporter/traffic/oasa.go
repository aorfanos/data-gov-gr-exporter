package traffic

import (
	"encoding/json"
	"log"

	"github.com/aorfanos/gov-gr-exporter/utils"
)

const apiEndpoint string = "https://data.gov.gr/api/v1/query/oasa_ridership"

type OasaRide []struct {
	Validations float64 `json:"dv_validations"`
	AgencyNum string `json:"dv_agency"`
	StationName string `json:"dv_platenum_station"`
	LoadTimestamp string `json:"load_dt"`
	Timestamp string `json:"date_hour"`
}

func GetMassTransitValidationInfo() *OasaRide {
	dateFrom, dateTo := utils.GenerateDateToFrom()
	body, err := utils.NewGovGrGetRequest(apiEndpoint+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var trafficInfo OasaRide
	if err := json.Unmarshal(body, &trafficInfo); err != nil {
		log.Fatal(err)
	}

	return &trafficInfo
}
