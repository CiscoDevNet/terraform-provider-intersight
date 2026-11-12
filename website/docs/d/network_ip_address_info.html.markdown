---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_ip_address_info"
description: |-
        An IP address record for a network element interface. Captures the address string, address family, prefix length, subnet mask, address origin, address state, and interface type for each discovered IPv4 or IPv6 address on the out-of-band or inband management interfaces of a network element.

---

# Data Source: intersight_network_ip_address_info
An IP address record for a network element interface. Captures the address string, address family, prefix length, subnet mask, address origin, address state, and interface type for each discovered IPv4 or IPv6 address on the out-of-band or inband management interfaces of a network element.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_ip_address_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `address`:(string) The IPv4 or IPv6 address string. Empty placeholder addresses are not stored; a network.IpAddressInfo record is only created when a non-empty address is present. 
* `address_family`:(string) The address family of this IP address (IPv4 or IPv6).* `` - The address family is not available or not applicable.* `IPv4` - The address is an IPv4 address.* `IPv6` - The address is an IPv6 address. 
* `address_origin`:(string) The origin of this IP address assignment. Normalized from platform-specific values to DHCP, SLAAC, Static, LinkLocal, or NotApplicable.* `` - The address origin is not available or not applicable.* `DHCP` - The IP address was assigned via DHCP or DHCPv6.* `SLAAC` - The IPv6 address was assigned via Stateless Address Autoconfiguration (SLAAC).* `Static` - The IP address is statically configured.* `LinkLocal` - The IPv6 address is a link-local address auto-derived from the interface MAC. 
* `address_state`:(string) The operational state of this IP address (Preferred, Deprecated, Failed). Set to NotApplicable for IPv4 addresses where per-address state is not applicable.* `` - The address state is not available or not applicable (e.g., for IPv4 addresses).* `Preferred` - The IP address is preferred and fully operational.* `Deprecated` - The IP address is deprecated (preferred lifetime has expired).* `Failed` - The IP address has failed (e.g., Duplicate Address Detection failure for IPv6). 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interface_type`:(string) The management interface type this address is associated with. OutOfBand for the OOB management interface; InBand for the inband management interface.* `` - The interface type is not available or not applicable.* `OutOfBand` - The IP address is associated with the out-of-band management interface.* `InBand` - The IP address is associated with the in-band management interface. 
* `ip_address_error`:(string) Any error or diagnostic message associated with this IP address. Empty when no error is reported. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `prefix_length`:(string) The prefix length of this IP address in bits. Used for IPv6 addresses. Set to zero for IPv4 addresses where SubnetMask is used instead. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subnet_mask`:(string) The subnet mask of this IP address in dotted-decimal notation. Used for IPv4 addresses. Empty for IPv6 addresses where PrefixLength is used instead. 
 
