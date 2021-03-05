---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group"
description: |-
  Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
---

# Data Source: intersight_storage_pure_protection_group
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_protection_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the protection Group. 
* `prefix`:(string) Prefix used for all generated snapshots from the protection group. 
* `replication_enabled`:(bool) Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. 
* `size`:(int) Overall size of all snapshots in the protection group, represented in bytes. 
* `snapshot_enabled`:(bool) Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. 
* `nr_source`:(string) Name of PureStorage array name on which the protection group is created. 
 