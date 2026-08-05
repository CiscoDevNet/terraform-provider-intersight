---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_inter_switch_port"
description: |-
        An Inter-Switch Port (ISP) represents the port which connects the two switches. In the context of Unified Edge, an ISP typically refers to the connection between two eCMCs in a chassis. This link provides redundancy for network connectivity.

---

# Data Source: intersight_ether_inter_switch_port
An Inter-Switch Port (ISP) represents the port which connects the two switches. In the context of Unified Edge, an ISP typically refers to the connection between two eCMCs in a chassis. This link provides redundancy for network connectivity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ether_inter_switch_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_vlan`:(string) Access VLAN for this port. 
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_fec`:(string) Administratively configured FEC mode for this port.* `` - Default value for FEC state.* `Not Supported` - FEC is not supported on this port.* `Disabled` - FEC is disabled on this port.* `Auto` - FEC mode is automatically negotiated between link partners.* `Cl74` - FEC is configured to use the IEEE Clause 74 (FireCode) standard.* `RS-IEEE(Cl108)` - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard.* `KP` - FEC is configured to use the KP (Backplane Ethernet) FEC mode. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port. 
* `allowed_vlans`:(string) Allowed VLANs on this port. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mac_address`:(string) Mac Address of a port in the Fabric Interconnect. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Operating mode of this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `native_vlan`:(string) Native VLAN for this port. 
* `oper_fec`:(string) Operational FEC mode for this port.* `` - Default value for FEC state.* `Not Supported` - FEC is not supported on this port.* `Disabled` - FEC is disabled on this port.* `Auto` - FEC mode is automatically negotiated between link partners.* `Cl74` - FEC is configured to use the IEEE Clause 74 (FireCode) standard.* `RS-IEEE(Cl108)` - FEC is configured to use the IEEE Clause 108 (Reed-Solomon) standard.* `KP` - FEC is configured to use the KP (Backplane Ethernet) FEC mode. 
* `oper_speed`:(string) Current Operational speed for this port. 
* `oper_state`:(string) Operational state of this port (enabled/disabled). 
* `oper_state_qual`:(string) Reason for this port's Operational state. 
* `oper_vlans`:(string) Operational VLANs on this port. 
* `peer_dn`:(string) PeerDn for ethernet physical port. 
* `port_id`:(int) Switch physical port identifier. 
* `port_name`:(string) Switch physical port name. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The role assigned to this port. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Speed of the inter switch link. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
 
