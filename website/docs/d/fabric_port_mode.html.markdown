---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_mode"
description: |-
        Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.

---

# Data Source: intersight_fabric_port_mode
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_port_mode.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `custom_mode`:(string) Custom Port Mode specified for the port range.* `FibreChannel` - Fibre Channel Port Types.* `BreakoutEthernet10G` - Breakout Ethernet 10G Port Type.* `BreakoutEthernet25G` - Breakout Ethernet 25G Port Type.* `BreakoutFibreChannel8G` - Breakout FibreChannel 8G Port Type.* `BreakoutFibreChannel16G` - Breakout FibreChannel 16G Port Type.* `BreakoutFibreChannel32G` - Breakout FibreChannel 32G Port Type. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id_end`:(int) Ending range of the Port Identifier. 
* `port_id_start`:(int) Starting range of the Port Identifier. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot Identifier of the switch. 
 
