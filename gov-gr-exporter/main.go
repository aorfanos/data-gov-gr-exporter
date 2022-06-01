package main

import (
	"log"
	"net/http"

	"github.com/aorfanos/gov-gr-exporter/property"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// trafficCollector := traffic.NewAthensTrafficCollector()
	ownersCollector := property.NewPropertyOwnersCollector()

	// prometheus.MustRegister(trafficCollector)
	prometheus.MustRegister(ownersCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":13211", nil))
}