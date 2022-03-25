---
subcategory: "top"
layout: "intersight"
page_title: "Intersight: intersight_top_system"
description: |-
        Root container for all UCSM / CIMC MOs.

---

# Data Source: intersight_top_system
Root container for all UCSM / CIMC MOs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_top_system.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ipv4_address`:(string) The IPv4 address of system. 
* `ipv6_address`:(string) The IPv6 address of system. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) The current mode of the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The admin configured name of the system. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `time_zone`:(string) The operational timezone of the system, empty indicates no timezone has been set specifically. 
 
