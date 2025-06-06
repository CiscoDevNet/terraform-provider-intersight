---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_dest_eth_port"
description: |-
        Configures Ethernet SPAN Destination Port for a given SPAN session.

---

# Data Source: intersight_fabric_span_dest_eth_port
Configures Ethernet SPAN Destination Port for a given SPAN session.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_dest_eth_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_speed`:(string) Admin speed of the SPAN Destination Port.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps.* `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fec`:(string) Forward error correction setting of the SPAN Destination Port.* `Auto` - Forward error correction option 'Auto'. Supported speeds are Auto, 1Gbps, 10Gbps, 25Gbps, 40Gbps and 100 Gbps.* `Cl91` - Forward error correction option 'cl91'. Supported speeds are 25Gbps and 100 Gbps.* `Cl74` - Forward error correction option 'cl74'. Supported speeds are 25Gbps.* `rs-cons16` - Forward error correction option \ rs-cons16\ . Supported speeds are 25Gbps.* `rs-ieee` - Forward error correction option \ rs-ieee\ . Supported speeds are 25Gbps.* `Off` - Turn off forward error correction. Supported speeds are 25Gbps and 100 Gbps. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot Identifier of the Switch Interface. 
 
