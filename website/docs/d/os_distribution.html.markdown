---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_distribution"
description: |-
  Intersight has the distribution details for all the Intersight supported OS
distributions. There will be a Distribution object for each supported OS.
---

# Data Source: intersight_os_distribution
Intersight has the distribution details for all the Intersight supported OS
distributions. There will be a Distribution object for each supported OS.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_distribution.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS distribution such as ESXi, CentOS. 
 