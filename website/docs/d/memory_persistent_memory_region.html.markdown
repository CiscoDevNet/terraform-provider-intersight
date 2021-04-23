---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_region"
description: |-
  Persistent Memory Region configured on the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory_persistent_memory_region
Persistent Memory Region configured on the Persistent Memory Modules on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_region.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_capacity`:(string) Free capacity in GiB of the Persistent Memory Region. 
* `health_state`:(string) Health state of the Persistent Memory Region. 
* `interleaved_set_id`:(string) ID of the Interleaved Set formed for this Persistent Memory Region. 
* `locater_ids`:(string) Set of locator IDs that are included in the Persistent Memory Region. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `persistent_memory_type`:(string) Persistent Memory type of the Persistent Memory Region. 
* `region_id`:(string) ID of the Persistent Memory Region. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `socket_id`:(string) Socket ID of the Persistent Memory Region. 
* `socket_memory_id`:(string) Socket Memory ID of the Persistent Memory Region. 
* `total_capacity`:(string) Total capacity in GiB of the Persistent Memory Region. 
 