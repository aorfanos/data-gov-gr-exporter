package energy

import (
	"fmt"

	"github.com/aorfanos/gov-gr-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type EnergyCollector struct {
	RenewableEnergyProductionMwh *prometheus.Desc
	EnergySystemLoadMwh *prometheus.Desc
	EnergyBalanceConsumptionMwh *prometheus.Desc
	EnergyBalanceConsumptionPercent *prometheus.Desc
}

func (c *EnergyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.RenewableEnergyProductionMwh
	ch <- c.EnergySystemLoadMwh
	ch <- c.EnergyBalanceConsumptionMwh
	ch <- c.EnergyBalanceConsumptionPercent
}

func NewRenewableEnergyCollector() *EnergyCollector {
	return &EnergyCollector{
		RenewableEnergyProductionMwh: prometheus.NewDesc("energy_renewable_production_count", "Count Mwh produced from renewable energy sources", []string{"timestamp_human"},
	prometheus.Labels{"unit": "Mwh"}),
	    EnergySystemLoadMwh: prometheus.NewDesc("energy_system_load_mwh", "Total system load in Mwh", []string{"timestamp_human"},
	prometheus.Labels{"unit": "Mwh"}),
	EnergyBalanceConsumptionMwh: prometheus.NewDesc("energy_balance_consumption_mwh", "Energy balance per fuel in Mwh", []string{"fuel_type","percentage","timestamp_human"}, nil),
	EnergyBalanceConsumptionPercent: prometheus.NewDesc("energy_balance_consumption_percent", "Energy balance percentage per fuel", []string{"fuel_type","consumption_mwh","timestamp_human"}, nil),

	}
}

func (c *EnergyCollector) Collect(ch chan<- prometheus.Metric) {
	renewableEnergyInfo := GetRenewableEnergyProduction("2022-05-25", "2022-05-31")
	energySystemLoad := GetEnergySystemLoad("2022-05-25", "2022-05-31")
	energyBalance := GetEnergyBalance("2022-05-25", "2022-05-31")

	for _, production := range *renewableEnergyInfo {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(production.Date), prometheus.MustNewConstMetric(
			c.RenewableEnergyProductionMwh, prometheus.GaugeValue,
			production.Production, production.Date))
	}

	for _, load := range *energySystemLoad {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(load.Date), prometheus.MustNewConstMetric(
			c.EnergySystemLoadMwh, prometheus.GaugeValue,
			load.Load, load.Date))
	}

	for _, balance := range *energyBalance {
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(balance.Date), prometheus.MustNewConstMetric(
			c.EnergyBalanceConsumptionMwh, prometheus.GaugeValue,
			balance.Energy, utils.GreekToLatin(balance.Fuel), fmt.Sprintf("%f", balance.Percentage), balance.Date))
		ch <- prometheus.NewMetricWithTimestamp(utils.DateStringToTimeZulu(balance.Date), prometheus.MustNewConstMetric(
			c.EnergyBalanceConsumptionPercent, prometheus.GaugeValue,
			balance.Percentage, utils.GreekToLatin(balance.Fuel), fmt.Sprintf("%f", balance.Energy), balance.Date))
	
	}
}