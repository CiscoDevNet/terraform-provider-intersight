---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_interface"
description: |-
        Details for the WAN Edge Device L3 interfaces (Portchannel, sub-interfaces, tunnels, SVIs, Routed ports).

---

# Data Source: intersight_catalystsdwan_interface
Details for the WAN Edge Device L3 interfaces (Portchannel, sub-interfaces, tunnels, SVIs, Routed ports).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) The administrative state of the WAN Edge Device interface.* `Unknown` - Administrative state is unknown.* `Up` - Administrative state is up.* `Down` - Administrative state is down. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) The distinguished name of the interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interface_name`:(string) The name of the interface on the WAN Edge Device. 
* `interface_type`:(string) The type of the interface (Portchannel, Sub-interfaces, tunnels, SVIs, Routed ports).* `Unknown` - Interface type is unknown.* `Loopback` - Interface type is loopback.* `Port Channel` - Interface type is port channel.* `Sub-Interface` - Interface type is sub-interface.* `Tunnel` - Interface type is tunnel.* `SVI` - Interface type is SVI (Switched Virtual Interface).* `Routed Port` - Interface type is routed port. 
* `ip_address`:(string) The IP address assigned to the WAN Edge Device interface. 
* `is_native_vlan`:(bool) Indicates whether the sub-interface's VLAN is configured as a native VLAN. 
* `mac_address`:(string) The MAC address of the interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) The operational state of the WAN Edge Device interface.* `Unknown` - Operational state of physical port or port channel or interface is unknown.* `Up` - Operational state of physical port or port channel or interface is up.* `Down` - Operational state of physical port or port channel or interface is down. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subnet`:(string) The subnet configured on the WAN Edge Device interface. 
* `system_ip`:(string) The system IP address of the WAN Edge Device. 
* `uuid`:(string) The UUID of the WAN Edge Device to which this interface belongs. 
* `vlan_id`:(int) The VLAN ID configured on the sub-interface. 
* `vpn_id`:(int) The VPN ID associated with this interface. 
 
