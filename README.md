# Greek OpenData exporter

An OpenMetrics exporter for datasets available via data.gov.gr

## Requirements

- A data.gov.gr [API key](https://www.data.gov.gr/token/)

## Deployment

### Standalone
Todo

### Docker
Todo

### Kubernetes
Todo


## Metrics

| Metric name              | Type  | Corresponding API Endpoint  | Collector |
|--------------------------|-------|-----------------------------|-----------|
| govgr_confiscations_per_region_count     | Gauge |    Confiscations    | property |
| govgr_owners_per_region_count | Gauge |   Owners per OTA  | property |
| govgr_owners_horizontal_per_region_count | Gauge |   Horizontal owners per OTA  | property |
| govgr_traffic_athens_avg_speed | Gauge |   Average speed per measuring device per road  | property |