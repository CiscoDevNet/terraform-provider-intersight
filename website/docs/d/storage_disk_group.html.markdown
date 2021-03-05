---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group"
description: |-
  Group of one or more Spans to configure virtual drive.
---

# Data Source: intersight_storage_disk_group
Group of one or more Spans to configure virtual drive.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_disk_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name to identity this disk group in the controller. 
* `raid_type`:(string) Raid level of the virtual drives in this diskgroup. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 