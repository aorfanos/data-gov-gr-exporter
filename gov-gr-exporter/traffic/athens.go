package traffic

import (
	"encoding/json"
	"log"

	"github.com/aorfanos/gov-gr-exporter/utils"
)

const ApiEndpoint string = "https://data.gov.gr/api/v1/query/road_traffic_attica"

type RoadInfo []struct {
	RoadName string `json:"road_name"`
	RoadInfo string `json:"road_info"`
	AverageSpeed float64 `json:"average_speed"`
	CountedCars float64 `json:"countedcars"`
	Timestamp string `json:"appprocesstime"`
	DeviceId string `json:"deviceid"`
}

func GetAthensTraffic(dateFrom, dateTo string) *RoadInfo {
	body, err := utils.NewGovGrGetRequest(ApiEndpoint+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var trafficInfo RoadInfo
	if err := json.Unmarshal(body, &trafficInfo); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(trafficInfo)

	return &trafficInfo
}
