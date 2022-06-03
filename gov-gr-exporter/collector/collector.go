package collector

import (
	"log"
	"os"
	"strings"

	"github.com/aorfanos/gov-gr-exporter/energy"
	"github.com/aorfanos/gov-gr-exporter/property"
	"github.com/aorfanos/gov-gr-exporter/sailing"
	"github.com/aorfanos/gov-gr-exporter/traffic"
	"github.com/prometheus/client_golang/prometheus"
)

func EnableSelectedCollectors(collectors []string) {
	for _, collector := range collectors {
		switch collector {
		case "traffic":
			trafficCollector := traffic.NewAthensTrafficCollector()
			prometheus.MustRegister(trafficCollector)
		case "property":
			ownersCollector := property.NewPropertyOwnersCollector()
			prometheus.MustRegister(ownersCollector)
		case "energy":
			energyCollector := energy.NewRenewableEnergyCollector()
			prometheus.MustRegister(energyCollector)
		case "sailing":
			sailingCollector := sailing.NewSailingCollector()
			prometheus.MustRegister(sailingCollector)
		}
	}
}

func ParseCollectorFlag(input string) []string {
	collectors := strings.Split(input, ",")
	log.Printf("Enabling collectors: %s", input)
	return collectors
}

func FlagDefault() string {
	const key = "GOV_GR_COLLECTORS_ENABLE"
	const value = "traffic,property,energy,sailing"
	
	if os.Getenv(key) == "" {
		log.Printf("No collectors enabled by user config. Using default collectors: %s", value)
		os.Setenv(key, value)
		return value
	}
	
	return os.Getenv(key)
}