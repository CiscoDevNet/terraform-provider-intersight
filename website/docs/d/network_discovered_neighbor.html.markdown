---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_discovered_neighbor"
description: |-
        L2 neighbor (LLDP and CDP) available on the switch.

---

# Data Source: intersight_network_discovered_neighbor
L2 neighbor (LLDP and CDP) available on the switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_discovered_neighbor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Device identifier of the switch. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `management_address`:(string) The management address of the switch. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `neighbor_device_id`:(string) Device identifier of the neighboring device. 
* `neighbor_interface`:(string) Neighboring device interface. 
* `neighbor_port_id`:(string) PortID of the neighbor device configured. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_port_id`:(string) Name of the local interface. 
 
