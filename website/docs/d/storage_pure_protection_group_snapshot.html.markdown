---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group_snapshot"
description: |-
  Protection group snapshot entity in Pure protection group.
---

# Data Source: intersight_storage_pure_protection_group_snapshot
Protection group snapshot entity in Pure protection group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_protection_group_snapshot.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string) Protection group snapshot creation time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Protection group snapshot name which represents point-in-time copy of all members in protection group. 
* `size`:(int) Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. 
* `nr_source`:(string) Source protection group name on which the snapshot is created. 
 