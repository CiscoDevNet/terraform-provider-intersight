---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_iwo_license_count"
description: |-
  Customer operation object to request reservation code.
---

# Data Source: intersight_license_iwo_license_count
Customer operation object to request reservation code.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_iwo_license_count.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `vm_license_count`:(int) The total number of devices claimed in the Intersight account. 
 