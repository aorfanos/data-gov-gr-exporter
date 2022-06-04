package traffic

import (
	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type TrafficCollector struct {
	AvgSpeed *prometheus.Desc
	CarCount *prometheus.Desc
	MassTransitValidations *prometheus.Desc
}

func (c *TrafficCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.AvgSpeed
	ch <- c.CarCount
	ch <- c.MassTransitValidations
}

func NewTrafficCollector() *TrafficCollector {
	return &TrafficCollector{
		AvgSpeed: prometheus.NewDesc("govgr_traffic_athens_avg_speed", "Average speed of Athens traffic", []string{"road_name","road_info","timestamp_human","device_id"}, prometheus.Labels{"unit": "km/h"}),
		CarCount: prometheus.NewDesc("govgr_traffic_athens_car_count", "Car count of Athens traffic", []string{"road_name","road_info","timestamp_human","device_id"}, nil),
		MassTransitValidations: prometheus.NewDesc("govgr_traffic_mass_transit_ticket_validations_count", "Ticket validations for Athens mass transit/public transport system", []string{"agency","station","timestamp_human"}, nil),
	}
}

func (c *TrafficCollector) Collect(ch chan<- prometheus.Metric) {
	trafficInfo := GetAthensTraffic()
	busTrafficInfo := GetMassTransitValidationInfo()

	for _, cars := range *trafficInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(cars.Timestamp), prometheus.MustNewConstMetric(
			c.AvgSpeed, prometheus.GaugeValue,
			cars.AverageSpeed, utils.GreekToLatin(cars.RoadName), utils.GreekToLatin(cars.RoadInfo), cars.Timestamp, cars.DeviceId,))
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(cars.Timestamp),prometheus.MustNewConstMetric(
			c.CarCount, prometheus.GaugeValue, 
			cars.CountedCars, utils.GreekToLatin(cars.RoadName), utils.GreekToLatin(cars.RoadInfo), cars.Timestamp, cars.DeviceId))
	}

	for _, buses := range *busTrafficInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(buses.Timestamp), prometheus.MustNewConstMetric(
			c.MassTransitValidations, prometheus.GaugeValue,
			buses.Validations, buses.AgencyNum, utils.GreekToLatin(buses.StationName), buses.Timestamp,))
		}
}