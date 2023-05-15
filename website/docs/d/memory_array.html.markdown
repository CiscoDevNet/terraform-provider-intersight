---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_array"
description: |-
        Holder housing multiple memory units.

---

# Data Source: intersight_memory_array
Holder housing multiple memory units.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_array.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `array_id`:(int) The instance number of the memory array. 
* `cpu_id`:(int) ID of the CPU that access this memory array. 
* `create_time`:(string) The time when this managed object was created. 
* `current_capacity`:(string) Current capacity of all the memory units on a server. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_correction`:(string) The primary hardware error detection or correction method supported by the memory array. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `max_capacity`:(string) Maximum capacity of all the memory units on a server. 
* `max_devices`:(string) The maximum number of slots or sockets available for memory devices in the memory array. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_power_state`:(string) The power state indicator of the memory array. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
