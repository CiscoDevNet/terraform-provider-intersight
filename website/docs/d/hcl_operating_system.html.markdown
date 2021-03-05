---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_operating_system"
description: |-
  Collection used to store operating system details.
---

# Data Source: intersight_hcl_operating_system
Collection used to store operating system details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_operating_system.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(string) Version of the Operating System. 
 