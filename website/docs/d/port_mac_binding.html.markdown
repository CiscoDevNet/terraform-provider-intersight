---
subcategory: "port"
layout: "intersight"
page_title: "Intersight: intersight_port_mac_binding"
description: |-
  Establishes relationship between the ports and connected end points based on LLDP TLVs.
---

# Data Source: intersight_port_mac_binding
Establishes relationship between the ports and connected end points based on LLDP TLVs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_port_mac_binding.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Aggregate Port ID of the local Switch Interface. 
* `chassis_id`:(int) Chassis/FEX device idetifier that is local to a cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mac`:(string) Device ID value that is advertised and available as a part of LLDP TLV. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port ID of the local Switch Interface. 
* `port_mac`:(string) Port ID value that is advertised and available as a part of LLDP TLV. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot ID of the local Switch slot Interface. 
* `switch_id`:(int) Switch Identifier that is local to a cluster. 
 