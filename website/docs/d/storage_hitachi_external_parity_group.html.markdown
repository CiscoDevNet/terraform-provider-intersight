---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_external_parity_group"
description: |-
        A external parity group in Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_external_parity_group
A external parity group in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_external_parity_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocatable_open_volume_capacity`:(int) From among the open volumes in the external parity group, the total capacity of volumes to which paths can be allocated (KB). 
* `allocated_open_volume_capacity`:(int) From among the open volumes in the external parity group, the total capacity of volumes to which paths are allocated (KB). 
* `available_volume_capacity`:(int) Available capacity of the external parity group, represented in bytes. 
* `clpr_id`:(int) Number of CLPR to which the external parity group belongs. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `emulation_type`:(string) Emulation type of the external parity group. 
* `external_product_id`:(string) Storage system that is connected using the external storage connection functionality of Universal Volume Manager. 
* `largest_available_capacity`:(int) Maximum capacity of the non-volume areas in the external parity group (KB). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) External parity group number. 
* `reserved_open_volume_capacity`:(int) From among the open volumes in the external parity group, the total capacity of volumes which are reserved (KB). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_open_volume_capacity`:(int) Total volume capacity of the open volumes in the external parity group (KB). 
* `unallocated_open_volume_capacity`:(int) From among the open volumes in the external parity group, the total capacity of volumes to which paths are not allocated (KB). 
* `used_capacity_rate`:(int) Usage rate of the external parity group. 
 
