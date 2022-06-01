package energy

import (
	"encoding/json"
	"log"

	"github.com/aorfanos/gov-gr-exporter/utils"
)

const ApiEndpointRenewables string = "https://data.gov.gr/api/v1/query/admie_realtimescadares"
const ApiEndpointSysLoad string = "https://data.gov.gr/api/v1/query/admie_realtimescadasystemload"
const ApiEndpointEnergyBalance string = "https://data.gov.gr/api/v1/query/admie_dailyenergybalanceanalysis"

type RenewableEnergy []struct {
	Production float64 `json:"energy_mwh"`
	Date string `json:"date"`
}

type SystemLoad []struct {
	Load float64 `json:"energy_mwh"`
	Date string `json:"date"`
}

type EnergyBalance []struct {
	Energy float64 `json:"energy_mwh"`
	Percentage float64 `json:"percentage"`
	Date string `json:"date"`
	Fuel string `json:"fuel"`
}

func GetRenewableEnergyProduction(dateFrom, dateTo string) *RenewableEnergy {
	body, err := utils.NewGovGrGetRequest(ApiEndpointRenewables+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var energy RenewableEnergy
	if err := json.Unmarshal(body, &energy); err != nil {
		log.Fatal(err)
	}

	return &energy
}

func GetEnergySystemLoad(dateFrom, dateTo string) *SystemLoad {
	body, err := utils.NewGovGrGetRequest(ApiEndpointSysLoad+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var energy SystemLoad
	if err := json.Unmarshal(body, &energy); err != nil {
		log.Fatal(err)
	}

	return &energy
}

func GetEnergyBalance(dateFrom, dateTo string) *EnergyBalance {
	body, err := utils.NewGovGrGetRequest(ApiEndpointEnergyBalance+"?date_from="+dateFrom+"&date_to="+dateTo)
	if err != nil {
		log.Fatal(err)
	}

	var energy EnergyBalance
	if err := json.Unmarshal(body, &energy); err != nil {
		log.Fatal(err)
	}

	return &energy
}
