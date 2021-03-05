---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_universe"
description: |-
  Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.
---

# Data Source: intersight_ippool_universe
Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_universe.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 