---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_port_channel"
description: |-
        Details for the WAN Edge Device port channels.

---

# Data Source: intersight_catalystsdwan_port_channel
Details for the WAN Edge Device port channels.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_port_channel.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) The administrative state of the WAN Edge Device port channel.* `Unknown` - Administrative state is unknown.* `Up` - Administrative state is up.* `Down` - Administrative state is down. 
* `allowed_vlan_ids`:(string) The allowed VLAN IDs configured on the WAN Edge Device port channel. 
* `create_time`:(string) The time when this managed object was created. 
* `default_vlan_id`:(int) The default VLAN ID configured on the WAN Edge Device port channel. 
* `dn`:(string) The distinguished name of the port channel. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Portchannel mode on WAN Edge Device will be 'access' or 'trunk' when switchport mode is enabled; otherwise, the mode is 'routed'.* `Routed` - Port mode is routed. This mode will be set when switchport mode is disabled regardless of the port configuration.* `Access` - Port mode is access. This mode will be set when switchport mode is enabled and the port is configured as an access port.* `Trunk` - Port mode is trunk. This mode will be set when switchport mode is enabled and the port is configured as a trunk port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The MTU of the WAN Edge Device port channel. 
* `oper_speed`:(string) The operational speed (in Mbps) of the WAN Edge Device port channel. 
* `oper_state`:(string) The operational state of the WAN Edge Device port channel.* `Unknown` - Operational state of physical port or port channel or interface is unknown.* `Up` - Operational state of physical port or port channel or interface is up.* `Down` - Operational state of physical port or port channel or interface is down. 
* `port_channel_name`:(string) Port channel name of the WAN Edge Device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `system_ip`:(string) The IP address of the WAN Edge Device. 
* `uuid`:(string) The UUID of the WAN Edge Device to which this port channel belongs. 
 
