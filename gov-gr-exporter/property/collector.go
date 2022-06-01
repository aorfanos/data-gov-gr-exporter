package property

import (
	"strconv"

	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type PropertyOwnersCollector struct {
	OwnerCount *prometheus.Desc
	HorizontalOwnerCount *prometheus.Desc
	ConfiscationsCount *prometheus.Desc
}

func (c *PropertyOwnersCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OwnerCount
	ch <- c.HorizontalOwnerCount
	ch <- c.ConfiscationsCount
}

func NewPropertyOwnersCollector() *PropertyOwnersCollector {
	return &PropertyOwnersCollector{
		OwnerCount: prometheus.NewDesc("govgr_owners_per_region_count", "Count of owners per OTA", []string{"ota_id","ota_full_name","ota_name_en","timestamp_human"}, nil),
		HorizontalOwnerCount: prometheus.NewDesc("govgr_horizontal_owners_per_region_count", "Count of horizontal owners per OTA", []string{"ota_id","ota_full_name","ota_name_en","timestamp_human"}, nil),
		ConfiscationsCount: prometheus.NewDesc("govgr_confiscations_per_region_count", "Confiscations per OTA", []string{"ota_id","ota_full_name","ota_name_en","timestamp_human"}, nil),
	}
}

func (c *PropertyOwnersCollector) Collect(ch chan<- prometheus.Metric) {
	ownerInfo := GetPropertyOwnersPerOta("2022-05-25", "2022-05-26")
	horizontalOwnerInfo := GetPropertyOwnersPerOta("2022-05-25", "2022-05-26")
	confiscations := GetConfiscationsPerOta("2022-05-29", "2022-05-31")

	for _, owners := range *ownerInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimePlainDate(owners.Date), prometheus.MustNewConstMetric(
			c.OwnerCount, prometheus.GaugeValue,
			owners.OwnersCount, strconv.Itoa(owners.OtaId), utils.GreekToLatin(owners.OtaFullName), owners.OtaNameEng, owners.Date))
	}
	
	for _, horzOwners := range *horizontalOwnerInfo {
		ch <- prometheus.NewMetricWithTimestamp(
			utils.DateStringToTimePlainDate(horzOwners.Date),
			prometheus.MustNewConstMetric(
				c.HorizontalOwnerCount, prometheus.GaugeValue,
				horzOwners.OwnersCount, strconv.Itoa(horzOwners.OtaId), utils.GreekToLatin(horzOwners.OtaFullName), horzOwners.OtaNameEng, horzOwners.Date),
			)
	}

	for _, confs := range *confiscations {
		ch <- prometheus.NewMetricWithTimestamp(
			utils.DateStringToTimePlainDate(confs.Date),
			prometheus.MustNewConstMetric(
				c.ConfiscationsCount, prometheus.GaugeValue,
				confs.Confiscations, strconv.Itoa(confs.OtaId), utils.GreekToLatin(confs.OtaFullName), confs.OtaNameEng, confs.Date),
			)
	}
}