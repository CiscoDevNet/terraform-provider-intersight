---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_app_catalog"
description: |-
  A catalog for managing application settings for HyperFlex cluster configuration service.
---

# Data Source: intersight_hyperflex_app_catalog
A catalog for managing application settings for HyperFlex cluster configuration service.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_app_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(string) The catalog version used in HyperFlex cluster configuration service. 
 