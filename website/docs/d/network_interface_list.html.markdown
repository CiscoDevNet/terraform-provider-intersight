---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_interface_list"
description: |-
        List of interfaces available on the switch to describe the available port inventory information.

---

# Data Source: intersight_network_interface_list
List of interfaces available on the switch to describe the available port inventory information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_interface_list.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state of the interface list. 
* `allowed_vlans`:(string) Allowed VLANs of the interface list. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the interface list. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `display_name`:(string) Display name of the interface list. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_address`:(string) IP address of the interface list. 
* `ip_subnet`:(int) IP subnet of the interface list. 
* `mac`:(string) MAC address of the interface list. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit of the interface list. 
* `name`:(string) Name of the interface list. 
* `oper_state`:(string) Operational state of the interface list. 
* `port_channel_id`:(int) Port channel id for port channel created on FI switch. 
* `port_sub_type`:(string) Interface types supported in Network device like Subinterfaces, Breakout Interfaces. 
* `port_type`:(string) Port type of interface list. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(string) Slot id of the interface list. 
* `speed`:(string) Port speed of the interface list. 
* `speed_group`:(string) Speed Group of the interface list. 
* `vlan`:(string) VLAN of the interface list. 
 
