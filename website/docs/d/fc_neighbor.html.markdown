---
subcategory: "fc"
layout: "intersight"
page_title: "Intersight: intersight_fc_neighbor"
description: |-
        Concrete class for switch and N port virtualization neighbors present in various interfaces of a switch.

---

# Data Source: intersight_fc_neighbor
Concrete class for switch and N port virtualization neighbors present in various interfaces of a switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fc_neighbor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `peer_device_capability`:(string) This field defines if neighbor is a switch or an NPV device.* `Switch` - Switch type neighbors of an interface.* `NPV` - N Port Virtualization neighbors of an interface. 
* `peer_interface`:(string) Interface through which the relationship is established. 
* `peer_ip_address`:(string) IP address of the peer switch. 
* `peer_switch_name`:(string) Device Id of the neighbor switch. 
* `peer_wwn`:(string) World Wide Name of the neighbor switch. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
