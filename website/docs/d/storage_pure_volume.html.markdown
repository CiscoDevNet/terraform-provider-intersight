---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume"
description: |-
  A volume entity in PureStorage FlashArray.
---

# Data Source: intersight_storage_pure_volume
A volume entity in PureStorage FlashArray.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_volume.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created`:(string) Creation time of the volume. 
* `description`:(string) Short description about the volume. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `serial`:(string) Serial number of the volume. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `nr_source`:(string) Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot. 
 