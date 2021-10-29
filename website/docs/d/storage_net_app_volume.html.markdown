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
To access the ith object of the results obtained, use `data.intersight_storage_net_app_volume.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `autosize_mode`:(string) The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink.* `off` - The volume will not grow or shrink in size in response to the amount of used space.* `grow` - The volume will automatically grow when used space in the volume is above the grow threshold.* `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Storage container's creation time. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `export_policy_name`:(string) The name of the Export Policy. 
* `key`:(string) Unique identifier of NetApp Volume across data center. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_policy_name`:(string) The name of the Snapshot Policy. 
* `snapshot_policy_uuid`:(string) The UUID of the Snapshot Policy. 
* `snapshot_utilized_capacity`:(int) The total space used by Snapshot copies in the volume represented in bytes. 
* `state`:(string) The current state of a NetApp volume.* `offline` - Read and write access to the volume is not allowed.* `online` - Read and write access to the volume is allowed.* `error` - Storage volume state of error type.* `mixed` - The constituents of a FlexGroup volume are not all in the same state. 
* `type`:(string) NetApp volume type. The volume type can be Read-write, Data-protection, or Load-sharing.* `data-protection` - Prevents modification of the data on the Volume.* `read-write` - Data on the Volume can be modified.* `load-sharing` - The volume type is Load Sharing DP. 
* `uuid`:(string) Universally unique identifier of a NetApp Volume. 
 