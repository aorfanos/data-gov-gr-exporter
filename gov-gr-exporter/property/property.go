// Ota || OTA == Οργανισμός Τοπικής Αυτοδιοίκησης
package property

import (
	"encoding/json"
	"log"

	"github.com/aorfanos/gov-gr-exporter/utils"
)

const ApiEndpointOwnersPerOta string = "https://data.gov.gr/api/v1/query/ktm_owners"
const ApiEndpointHorizontalOwners string = "https://data.gov.gr/api/v1/query/ktm_hplots"
const ApiEndpointConfiscations string = "https://data.gov.gr/api/v1/query/ktm_confs"

type PropertyOwnersPerMunicipality []struct {
	OtaId int `json:"ota_id"`
	OwnersCount float64 `json:"owners"`
	OtaFullName string `json:"ota_full_name"`
	OtaNameEng string `json:"ota_name_en"` // as provided by data.gov.gr
	Date string `json:"date"`
}

type PropertyConfiscationsPerMunicipality []struct {
	OtaId int `json:"ota_id"`
	Confiscations float64 `json:"confiscations"`
	OtaFullName string `json:"ota_full_name"`
	OtaNameEng string `json:"ota_name_en"` // as provided by data.gov.gr
	Date string `json:"date"`
}

func GetPropertyOwnersPerOta(dateFrom, dateTo string) *PropertyOwnersPerMunicipality {
	body, err := utils.NewGovGrGetRequest(ApiEndpointOwnersPerOta+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var owners PropertyOwnersPerMunicipality
	if err := json.Unmarshal(body, &owners); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(owners)

	return &owners
}

func GetHorizontalPropertyOwnersPerOta(dateFrom, dateTo string) *PropertyOwnersPerMunicipality {
	body, err := utils.NewGovGrGetRequest(ApiEndpointHorizontalOwners+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var owners PropertyOwnersPerMunicipality
	if err := json.Unmarshal(body, &owners); err != nil {
		log.Fatal(err)
	}

	return &owners
}

func GetConfiscationsPerOta(dateFrom, dateTo string) *PropertyConfiscationsPerMunicipality {
	body, err := utils.NewGovGrGetRequest(ApiEndpointConfiscations+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var confiscations PropertyConfiscationsPerMunicipality
	if err := json.Unmarshal(body, &confiscations); err != nil {
		log.Fatal(err)
	}

	return &confiscations
}