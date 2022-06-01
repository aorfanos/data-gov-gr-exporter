package traffic

import (
	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type AthensTrafficCollector struct {
	AvgSpeed *prometheus.Desc
	CarCount *prometheus.Desc
}

func (c *AthensTrafficCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.AvgSpeed
	ch <- c.CarCount
}

func NewAthensTrafficCollector() *AthensTrafficCollector {
	return &AthensTrafficCollector{
		AvgSpeed: prometheus.NewDesc("govgr_traffic_athens_avg_speed", "Average speed of Athens traffic", []string{"road_name","road_info","timestamp_human","device_id"}, prometheus.Labels{"unit": "km/h"}),
		CarCount: prometheus.NewDesc("govgr_traffic_athens_car_count", "Car count of Athens traffic", []string{"road_name","road_info","timestamp_human","device_id"}, nil),
	}
}

func (c *AthensTrafficCollector) Collect(ch chan<- prometheus.Metric) {
	trafficInfo := GetAthensTraffic("2022-05-31", "2022-06-01")

	for _, cars := range *trafficInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(cars.Timestamp), prometheus.MustNewConstMetric(
			c.AvgSpeed, prometheus.GaugeValue,
			cars.AverageSpeed, utils.GreekToLatin(cars.RoadName), utils.GreekToLatin(cars.RoadInfo), cars.Timestamp, cars.DeviceId,))
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(cars.Timestamp),prometheus.MustNewConstMetric(
			c.CarCount, prometheus.GaugeValue, 
			cars.CountedCars, utils.GreekToLatin(cars.RoadName), utils.GreekToLatin(cars.RoadInfo), cars.Timestamp, cars.DeviceId))
	}
	
}