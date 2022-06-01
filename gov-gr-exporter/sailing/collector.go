package sailing

import (
	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type SailingCollector struct {
	PassengerCount *prometheus.Desc
	VehicleCount *prometheus.Desc
}

func (c *SailingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.PassengerCount
	ch <- c.VehicleCount
}

func NewSailingCollector() *SailingCollector {
	return &SailingCollector{
		PassengerCount: prometheus.NewDesc("govgr_sailing_traffic_passenger_count", "Passenger count for given trip",
		[]string{
			"arrival_port",
			"arrival_port_name",
			"departure_port",
			"departure_port_name",
			"route_code",
			"route_codes_names",
			"timestamp_human",
		}, nil),
		VehicleCount: prometheus.NewDesc("govgr_sailing_traffic_vehicle_count", "Vehicle count for given trip",
		[]string{
			"arrival_port",
			"arrival_port_name",
			"departure_port",
			"departure_port_name",
			"route_code",
			"route_codes_names",
			"timestamp_human",
		}, nil),

	}
}

func (c *SailingCollector) Collect(ch chan<- prometheus.Metric) {
	sailingTraffic := GetSailingTraffic()

	for _, traffic := range *sailingTraffic {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimePlainDate(traffic.Date),
		prometheus.MustNewConstMetric(
			c.PassengerCount, prometheus.GaugeValue,
			traffic.PassengerCount,
			traffic.ArrivalPort,
			utils.GreekToLatin(traffic.ArrivalPortName),
			traffic.DeparturePort,
			utils.GreekToLatin(traffic.DeparturePortName),
			traffic.RouteCode,
			utils.GreekToLatin(traffic.RouteCodeName),
			traffic.Date,
			))

		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimePlainDate(traffic.Date),
		prometheus.MustNewConstMetric(
			c.VehicleCount, prometheus.GaugeValue,
			traffic.VehicleCount,
			traffic.ArrivalPort,
			utils.GreekToLatin(traffic.ArrivalPortName),
			traffic.DeparturePort,
			utils.GreekToLatin(traffic.DeparturePortName),
			traffic.RouteCode,
			utils.GreekToLatin(traffic.RouteCodeName),
			traffic.Date,
			))
	}
}

// "arrivalport": "PAA",
// "arrivalportname": "Ψαρά",
// "date": "2022-05-29",
// "departureport": "JKH",
// "departureportname": "Χίος",
// "passengercount": 8,
// "routecode": "MJTJKHINOPAAPIR",
// "routecodenames": "Μυτιλήνη-Χίος-Οινούσσες-Ψαρά-Πειραιάς",