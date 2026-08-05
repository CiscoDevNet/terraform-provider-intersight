---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_physical_port"
description: |-
        Physical ports belonging to a WAN Edge Device.

---

# Data Source: intersight_catalystsdwan_physical_port
Physical ports belonging to a WAN Edge Device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_physical_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_speed`:(string) Admin speed of the physical port. 
* `admin_state`:(string) Administrative state of the physical port.* `Unknown` - Administrative state is unknown.* `Up` - Administrative state is up.* `Down` - Administrative state is down. 
* `allowed_vlan_ids`:(string) List of configured allowed VLAN IDs. 
* `auto_negotiation`:(bool) Indicates if auto-negotiation is enabled. 
* `create_time`:(string) The time when this managed object was created. 
* `default_vlan_id`:(int) Default or Native VLAN configured on the port. 
* `dn`:(string) The distinguished name of the physical port. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fec_enabled`:(bool) Indicates if FEC is enabled on the physical port. 
* `mac_address`:(string) MAC address of the physical port. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Configured physical port mode. Value will be 'access' or 'trunk' when configured as a switch port; 'backplane' for internal backplane ports and 'routed' all other ports.* `Routed` - Port mode is routed. This mode will be set when switchport mode is disabled regardless of the port configuration.* `Access` - Port mode is access. This mode will be set when switchport mode is enabled and the port is configured as an access port.* `Trunk` - Port mode is trunk. This mode will be set when switchport mode is enabled and the port is configured as a trunk port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) MTU configured on the physical port. 
* `negotiated_speed`:(string) Negotiated speed of the physical port. 
* `oper_state`:(string) Operational state of the physical port.* `Unknown` - Operational state of physical port or port channel or interface is unknown.* `Up` - Operational state of physical port or port channel or interface is up.* `Down` - Operational state of physical port or port channel or interface is down. 
* `port_channel_id`:(int) Port channel ID if the physical port is part of a port channel. 
* `port_channel_name`:(string) Port channel name if the physical port is part of a port channel. 
* `port_name`:(string) Physical port name of the WAN Edge Device. 
* `port_type`:(string) Type of the physical port (e.g., Backplane and frontpanel).* `Front Panel` - The physical ports that are located on the front panel of Catalyst SDWAN devices.* `Backplane` - The internal backplane ports on Catalyst SDWAN devices. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `system_ip`:(string) The IP address of the WAN Edge Device. 
* `uuid`:(string) The UUID of the WAN Edge Device to which this physical port belongs. 
 
