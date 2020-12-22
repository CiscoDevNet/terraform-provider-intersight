---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_uplink_role"
description: |-
  Configuration object sent by user to create a fc uplink port.
---

# Data Source: intersight_fabric_fc_uplink_role
Configuration object sent by user to create a fc uplink port.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `8Gbps` - Admin configurable speed 8Gbps.* `16Gbps` - Admin configurable speed 16Gbps.* `32Gbps` - Admin configurable speed 32Gbps. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `fill_pattern`:(string) Fill pattern to differentiate the configs in NPIV.* `Idle` - Fc Fill Pattern type Idle.* `Arbff` - Fc Fill Pattern type Arbff. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
* `vsan_id`:(int) Virtual San Identifier associated to the FC port. 
