---
subcategory: "management"
layout: "intersight"
page_title: "Intersight: intersight_management_interface"
description: |-
  Interface that provides access to the management controller.
---

# Data Source: intersight_management_interface
Interface that provides access to the management controller.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_management_interface.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gateway`:(string) Default gateway for the interface. 
* `host_name`:(string) Hostname configured for the interface. 
* `ip_address`:(string) IP address of the interface. 
* `ipv4_address`:(string) IPv4 address of the interface. 
* `ipv4_gateway`:(string) IPv4 default gateway for the interface. 
* `ipv4_mask`:(string) IPv4 Netmask for the interface. 
* `ipv6_address`:(string) IPv6 address of the interface. 
* `ipv6_gateway`:(string) IPv6 default gateway for the interface. 
* `ipv6_prefix`:(int) IPv6 prefix for the interface. 
* `mac_address`:(string) MAC address configured for the interface. 
* `mask`:(string) Netmask for the interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) Switch Id connected to the interface. 
* `uem_conn_status`:(string) The event channel connection status for the interface. 
* `virtual_host_name`:(string) Virtual hostname configured for the interface in case of clustered environment. 
* `vlan_id`:(int) VlanId configured for the interface. 
 