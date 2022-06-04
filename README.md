[![Package](https://github.com/aorfanos/data-gov-gr-exporter/actions/workflows/package.yaml/badge.svg)](https://github.com/aorfanos/data-gov-gr-exporter/actions/workflows/package.yaml)
[![Dockerhub](https://github.com/aorfanos/data-gov-gr-exporter/actions/workflows/dockerhub.yaml/badge.svg)](https://github.com/aorfanos/data-gov-gr-exporter/actions/workflows/dockerhub.yaml)
# Greek OpenData exporter

A Prometheus/OpenMetrics exporter for datasets available via data.gov.gr

## Requirements

- A data.gov.gr [API key](https://www.data.gov.gr/token/)

## Usage

```shell
Flags:
  --help        Show context-sensitive help (also try --help-long and --help-man).
  --port=13211  Port to listen on.
  --collectors="traffic,property,energy,sailing"  
                Comma separated list of collectors to enable.
```

## Deployment

### Standalone

Grab one of the binaries in https://github.com/aorfanos/data-gov-gr-exporter/releases/latest

### Docker

- Pull the image from Dockerhub
```shell
docker pull saikolab/data-gov-gr-exporter:latest
```

- Run it
```shell
docker run \
    --publish 13211:13211 \
    -e"GOV_GR_API_KEY=<YOUR-API-KEY>" \
    --name gov-exporter \
    -d \
    saikolab/data-gov-gr-exporter:latest
```

### Kubernetes
Todo


## Metrics

| Metric name              | Type  | Corresponding API Endpoint  | Collector |
|--------------------------|-------|-----------------------------|-----------|
| govgr_confiscations_per_region_count     | Gauge |    Confiscations    | property |
| govgr_owners_per_region_count | Gauge |   Owners per OTA  | property |
| govgr_owners_horizontal_per_region_count | Gauge |   Horizontal owners per OTA  | property |
| govgr_traffic_athens_avg_speed | Gauge |   Average speed per measuring device per road  | property |
| govgr_energy_renewable_production_count | Gauge |   Count Mwh produced from renewable energy sources  | energy |
| govgr_energy_system_load_mwh | Gauge |   Total system load in Mwh  | energy |
| govgr_energy_balance_consumption_mwh | Gauge |   Energy balance per used fuel in Mwh  | energy |
| govgr_energy_balance_consumption_percent | Gauge |   Energy balance percentage per fuel  | energy |
| govgr_sailing_traffic_passenger_count | Gauge |   Passenger count for given trip  | sailing |
| govgr_sailing_traffic_vehicle_count | Gauge |   Vehicle count for given trip  | sailing |
| govgr_traffic_athens_avg_speed | Gauge |   Average speed of Athens traffic  per road per direction | sailing |
| govgr_traffic_athens_car_count | Gauge |   Car count of Athens traffic per road  per direction | sailing |