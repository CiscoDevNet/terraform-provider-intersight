---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_volume"
description: |-
  NetApp volume are data containers that enable you to partition and manage your data.
---

# Data Source: intersight_storage_net_app_volume
NetApp volume are data containers that enable you to partition and manage your data.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_volume.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `autosize_mode`:(string) The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink.* `off` - The volume will not grow or shrink in size in response to the amount of used space.* `grow` - The volume will automatically grow when used space in the volume is above the grow threshold.* `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space. 
* `created_time`:(string) Storage container's creation time. 
* `export_policy_name`:(string) Name of Export Policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `snapshot_policy_name`:(string) Name of Snapshot Policy. 
* `snapshot_policy_uuid`:(string) Name of Snapshot Policy. 
* `snapshot_utilized_capacity`:(int) The total space used by Snapshot copies in the volume represented in bytes. 
* `state`:(string) The current state of a NetApp volume.* `offline` - Read and write access to the volume is not allowed.* `online` - Read and write access to the volume is allowed.* `error` - Storage volume state of error type.* `mixed` - The constituents of a FlexGroup volume are not all in the same state. 
* `type`:(string) NetApp volume type. The volume type can be Read-write or Data-protection, Load-sharing, or Data-cache.* `data-protection` - Prevents modification of the data on the Volume.* `read-write` - Data on the Volume can be modified.* `load-sharing` - Load Sharing. 
* `uuid`:(string) UUID of NetApp Volume. 
 