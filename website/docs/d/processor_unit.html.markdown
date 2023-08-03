---
subcategory: "processor"
layout: "intersight"
page_title: "Intersight: intersight_processor_unit"
description: |-
        The CPU present on a server.

---

# Data Source: intersight_processor_unit
The CPU present on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_processor_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `architecture`:(string) The architecture of the installed processor. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field displays the description of the processor. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_platform_supported`:(bool) This field indicates whether the processor is supported on the server or not. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_cores`:(int) The number of cores present in a given processor. 
* `num_cores_enabled`:(string) The number of enabled cores in the installed processor. 
* `num_threads`:(string) The maximum number of threads available in the installed processor. 
* `oper_power_state`:(string) The power state of the processor. 
* `oper_state`:(string) The health indicator of the processor, 'OK' indicates the processor is operatinal. 
* `operability`:(string) Operability state of the central processing unit. 
* `part_number`:(string) This field displays the part number of the of the processor. 
* `pid`:(string) This field displays the product ID of the processor. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `processor_id`:(int) The ID number of a given processor. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `socket_designation`:(string) The socket ID of the installed processor. 
* `speed`:(float) The maximum speed of the installed processor in GHz. 
* `stepping`:(string) The CPU stepping of the installed processor. 
* `thermal`:(string) The temperature of the processor in centigrade. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
