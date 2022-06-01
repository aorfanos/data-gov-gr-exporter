package property

import (
	"strconv"

	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type PropertyOwnersCollector struct {
	OwnerCount *prometheus.Desc
}

func (c *PropertyOwnersCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OwnerCount
}

func NewPropertyOwnersCollector() *PropertyOwnersCollector {
	return &PropertyOwnersCollector{
		OwnerCount: prometheus.NewDesc("owners_per_region_count", "Count of owners per OTA", []string{"ota_id","ota_full_name","ota_name_en","timestamp_human"}, nil),
	}
}

func (c *PropertyOwnersCollector) Collect(ch chan<- prometheus.Metric) {
	ownerInfo := GetPropertyOwnersPerOta("2022-05-25", "2022-05-26")

	for _, owners := range *ownerInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimePlainDate(owners.Date), prometheus.MustNewConstMetric(
			c.OwnerCount, prometheus.GaugeValue,
			owners.OwnersCount, strconv.Itoa(owners.OtaId), utils.GreekToLatin(owners.OtaFullName), owners.OtaNameEng, owners.Date))
	}
	
}