---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk_usage"
description: |-
  Has usage map between physical disks and virtual drives.
---

# Data Source: intersight_storage_physical_disk_usage
Has usage map between physical disks and virtual drives.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_physical_disk_usage.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_blocks`:(string) The number of blocks that are a part of the virtual drive. 
* `physical_drive`:(string) The physical disk for which the usage is reported. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `span`:(string) The span of the physical disk. 
* `starting_block`:(string) The starting block id of the virtual drive within the physical drive. 
* `state`:(string) The current state of the physical disk usage. 
* `virtual_drive`:(string) The virtual drive corresponding to the physical disk. 
 