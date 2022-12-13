---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_reservation"
description: |-
        The IP reservation object, used to hold reserved addresses.

---

# Data Source: intersight_ippool_reservation
The IP reservation object, used to hold reserved addresses.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_reservation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the allocation for the identity in the reservation either static or dynamic (i.e. via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) IP identity to be reserved. 
* `ip_type`:(string) Type of the IP address that needs to be reserved.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address to be reserved. 
* `ip_v6_address`:(string) IPv6 Address to be reserved. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vrf_moid`:(string) The moid of the Virtual Routing and Forwarding MO. 
 
