package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/aorfanos/gov-gr-exporter/collector"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port = kingpin.Flag("port", "Port to listen on.").Default("13211").Int()
	registerCollectors = kingpin.Flag("collectors", "Comma separated list of collectors to enable.").Default("traffic,property,energy,sailing").String()
)

func main() {
	kingpin.Parse()

	collector.EnableSelectedCollectors(collector.ParseCollectorFlag(*registerCollectors))

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting data.gov.gr exporter on port :%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}