---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_supervisor_card"
description: |-
        Concrete class for supervisor card.

---

# Data Source: intersight_network_supervisor_card
Concrete class for supervisor card.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_supervisor_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of the supervisor card. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hardware_version`:(string) The hardware version of the supervisor card. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the supervisor card like Supervisor Card-1. 
* `number_of_ports`:(int) The number of ports on the supervisor card. 
* `oper_reason`:(string) The health issue of the supervisor card. 
* `oper_state`:(string) The operational status of the supervisor card. 
* `part_number`:(string) The part number of the supervisor card. 
* `power_state`:(string) The power state of the supervisor card. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the supervisor card. 
* `supervisor_id`:(string) The identifier for the supervisor card. 
* `type`:(string) The type of the supervisor card. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
