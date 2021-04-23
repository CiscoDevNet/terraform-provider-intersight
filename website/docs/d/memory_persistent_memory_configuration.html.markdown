---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_configuration"
description: |-
  Persistent Memory configuration on all the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory_persistent_memory_configuration
Persistent Memory configuration on all the Persistent Memory Modules on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `memory_capacity`:(string) Memory capacity in GiB of a Persistent Memory configuration on a server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_modules`:(string) Number of Persistent Memory Modules on a server. 
* `num_of_regions`:(string) Number of Persistent Memory Regions on a server. 
* `persistent_memory_capacity`:(string) Persistent memory capacity in GiB of a Persistent Memory configuration on a server. 
* `reserved_capacity`:(string) Reserved capacity in GiB of a Persistent Memory configuration on a server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `security_state`:(string) Collective security state of all Persistent Memory modules on a server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_capacity`:(string) Total capacity in GiB of a Persistent Memory configuration on a server. 
 