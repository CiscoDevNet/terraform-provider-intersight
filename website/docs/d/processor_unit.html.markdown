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
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_cores`:(int) The number of cores present in a given processor. 
* `num_cores_enabled`:(string) The number of enabled cores in the installed processor. 
* `num_threads`:(string) The maximum number of threads available in the installed processor. 
* `oper_power_state`:(string) The power state of the processor. 
* `oper_state`:(string) The health indicator of the processor, 'OK' indicates the processor is operatinal. 
* `operability`:(string) Operability state of the central processing unit. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `processor_id`:(int) The ID number of a given processor. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `socket_designation`:(string) The socket ID of the installed processor. 
* `speed`:(float) The maximum speed of the installed processor in GHz. 
* `stepping`:(string) The CPU stepping of the installed processor. 
* `thermal`:(string) The temperature of the processor in centigrade. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 
