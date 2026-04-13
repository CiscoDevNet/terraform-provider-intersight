---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_net_flow_record"
description: |-
        The NetFlow record represents a single network flow, containing key attributes (source/destination address, protocol, tos) and statistics collected like packet count, byte count and timestamp.

---

# Data Source: intersight_fabric_net_flow_record
The NetFlow record represents a single network flow, containing key attributes (source/destination address, protocol, tos) and statistics collected like packet count, byte count and timestamp.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_net_flow_record.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Netflow record name. Must be a maximum of 63 characters, without spacing. 
* `record_type`:(string) Netflow Record Type IPv4, IPv6 and L2.* `Invalid` - Netflow record invlaid type.* `IPv4` - Netflow record type for IPv4 packet.* `IPv6` - Netflow record type for IPv6 packet.* `L2` - Netflow record type for L2 packet. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
