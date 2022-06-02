# Greek OpenData exporter

An OpenMetrics exporter for datasets available via data.gov.gr

## Requirements

- A data.gov.gr [API key](https://www.data.gov.gr/token/)

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