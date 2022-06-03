package collector

import (
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
	return collectors
}