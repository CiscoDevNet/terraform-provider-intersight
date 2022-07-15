---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_snapshot"
description: |-
        A snapshot entity in Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_snapshot
A snapshot entity in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `concordance_rate`:(int) Concordance rate for pairs. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_consistency_group`:(bool) Whether the pair was created in the consistency group mode (CTG mode). 
* `is_multistageable`:(bool) Whether the pair can be a multistage pair. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mu_number`:(int) MU number of the primary volume. 
* `name`:(string) Object ID of the pair for snapshot data. 
* `pvol_ldev_id`:(int) LDEV number of the primary volume. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_group_name`:(string) Name of the snapshot group that contains the pair for snapshot data. 
* `snapshot_pool_id`:(int) ID of the pool in which the snapshot data is created. 
* `split_time`:(string) Time when snapshot data was created. 
* `status`:(string) Pair status. Pair status changes according to the pair operation. 
* `svol_ldev_id`:(int) LDEV number of the secondary volume. 
 
