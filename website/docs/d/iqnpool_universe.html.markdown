---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_universe"
description: |-
  Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
---

# Data Source: intersight_iqnpool_universe
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iqnpool_universe.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 