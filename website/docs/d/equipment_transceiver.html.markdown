---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_transceiver"
description: |-
        Transceiver information on the chassis.

---

# Data Source: intersight_equipment_transceiver
Transceiver information on the chassis.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_transceiver.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Breakout port member in the Fabric Interconnect. 
* `cisco_extended_id_number`:(string) The cisco extended Id number state of the pluggable SFP. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interface_type`:(string) Interface type of transceiver copper or fiber. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `manufacturer_part_number`:(string) The manufacturer part number of the pluggable SFP. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `module_id`:(int) Fabric extender identifier. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the pluggable transceiver. 
* `oper_speed`:(string) Operational speed of the transceiver. 
* `oper_state`:(string) Operational state of the transceiver. 
* `oper_state_qual`:(string) Reason for this transceiver's operational state. 
* `port_id`:(int) Switch physical port identifier. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `status`:(string) Status of the pluggable SFP. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `type`:(string) The type of the transceiver. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
