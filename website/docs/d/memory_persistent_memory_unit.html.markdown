---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_unit"
description: |-
        Persistent Memory Module on a server.

---

# Data Source: intersight_memory_persistent_memory_unit
Persistent Memory Module on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) This represents the administrative state of the memory unit on a server. 
* `app_direct_capacity`:(string) AppDirect capacity in GiB of the Persistent Memory Module on a server. 
* `array_id`:(int) This represents the memory array to which the memory unit belongs to. 
* `bank`:(int) This represents the memory bank of the memory unit on a server. 
* `capacity`:(string) This represents the memory capacity in MiB of the memory unit on a server. 
* `clock`:(string) This represents the clock of the memory unit on a server. 
* `count_status`:(string) Count status of the Persistent Memory Module on a server. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) Firmware version of the firware running on the Persistent Memory Module on a server. 
* `form_factor`:(string) This represents the form factor of the memory unit on a server. 
* `frozen_status`:(string) Frozen status of the Persistent Memory Module on a server. 
* `health_state`:(string) Health state of the Persistent Memory Module on a server. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `latency`:(string) This represents the latency of the memory unit on a server. 
* `location`:(string) This represents the location of the memory unit on a server. 
* `lock_status`:(string) Lock status of the Persistent Memory Module on a server. 
* `memory_capacity`:(string) Memory capacity in GiB of the Persistent Memory Module on a server. 
* `memory_id`:(int) ID of the Persistent Memory Module on a server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_power_state`:(string) This represents the operational power state of the memory unit on a server. 
* `oper_state`:(string) This represents the operational state of the memory unit on a server. 
* `operability`:(string) This represents the operability of the memory unit on a server. 
* `persistent_memory_capacity`:(string) Persistent Memory capacity in GiB of the Persistent Memory Module on a server. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `reserved_capacity`:(string) Reserved capacity in GiB of the Persistent Memory Module on a server. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `security_status`:(string) Security status of the Persistent Memory Module on a server. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `set`:(int) This represents the set of the memory unit on a server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `socket_id`:(string) Socket ID of the Persistent Memory Module on a server. 
* `socket_memory_id`:(string) Socket Memory ID of the Persistent Memory Module on a server. 
* `speed`:(string) This represents the speed of the memory unit on a server. 
* `thermal`:(string) This represents the thremal state of the memory unit on a server. 
* `total_capacity`:(string) Total capacity in GiB of the Persistent Memory Module on a server. 
* `type`:(string) This represents the memory type of the memory unit on a server. 
* `uid`:(string) UID of the Persistent Memory Module on a server. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `visibility`:(string) This represents the visibility of the memory unit on a server. 
* `width`:(string) This represents the width of the memory unit on a server. 
 
