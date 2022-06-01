package main

import (
	"log"
	"net/http"

	"github.com/aorfanos/gov-gr-exporter/traffic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	trafficCollector := traffic.NewAthensTrafficCollector()
	prometheus.MustRegister(trafficCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":13211", nil))
}