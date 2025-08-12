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
To access the ith object of the results obtained, use `data.intersight_storage_pure_protection_group_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Protection group snapshot creation time. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `eradication_config`:(string) The configuration of eradication feature. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Protection group snapshot name which represents point-in-time copy of all members in protection group. 
* `pod`:(string) A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. 
* `snapshot_size`:(int) The size of the snapshot created. 
* `nr_source`:(string) Source protection group name on which the snapshot is created. 
* `total_provisioned`:(int) The overall size of the snapshot allocated by the storage array. 
 
