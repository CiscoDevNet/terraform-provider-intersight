---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_host_port"
description: |-
        Host Interface ports available on the I/O module or the Fabric Extender that facilitate connectivity between the Fabric Interconnect and the Cisco UCS B/C/X-Series servers.

---

# Data Source: intersight_ether_host_port
Host Interface ports available on the I/O module or the Fabric Extender that facilitate connectivity between the Fabric Interconnect and the Cisco UCS B/C/X-Series servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ether_host_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port. 
* `aggregate_port_id`:(int) Breakout port member in the fabric extender. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mac_address`:(string) Mac Address of a port in the Fabric Interconnect. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Operating mode of this port. 
* `module_id`:(int) Fabric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_speed`:(string) Current Operational speed for this port. 
* `oper_state`:(string) Operational state of this port (enabled/disabled). 
* `oper_state_qual`:(string) Reason for this port's Operational state. 
* `peer_dn`:(string) PeerDn for ethernet physical port. 
* `port_channel_id`:(int) Port channel id for port channel created on FI switch. 
* `port_id`:(int) Switch physical port identifier. 
* `port_type`:(string) Defines the transport type for this port (ethernet OR fc). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The role assigned to this port. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Host Port Speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `transceiver_type`:(string) Transceiver model attached to a port in the Fabric Interconnect. 
 
