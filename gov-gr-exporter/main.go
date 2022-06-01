package main

import (
	"log"
	"net/http"

	"github.com/aorfanos/gov-gr-exporter/energy"
	"github.com/aorfanos/gov-gr-exporter/property"
	"github.com/aorfanos/gov-gr-exporter/sailing"
	"github.com/aorfanos/gov-gr-exporter/traffic"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	trafficCollector := traffic.NewAthensTrafficCollector()
	ownersCollector := property.NewPropertyOwnersCollector()
	energyCollector := energy.NewRenewableEnergyCollector()
	sailingCollector := sailing.NewSailingCollector()


	prometheus.MustRegister(trafficCollector)
	prometheus.MustRegister(ownersCollector)
	prometheus.MustRegister(energyCollector)
	prometheus.MustRegister(sailingCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting data.gov.gr exporter on port :13211")
	log.Fatal(http.ListenAndServe(":13211", nil))
}