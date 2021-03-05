---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_health_check_execution_snapshot"
description: |-
  Health check execution snapshot of the HyperFlex cluster.
---

# Data Source: intersight_hyperflex_cluster_health_check_execution_snapshot
Health check execution snapshot of the HyperFlex cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_health_check_execution_snapshot.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `timestamp`:(string) Timestamp of the last health check execution on the HyperFlex cluster. 
 