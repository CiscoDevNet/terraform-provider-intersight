---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_operation"
description: |-
  PortOperation objects allows the user to alter the state of the port.
---

# Data Source: intersight_fabric_port_operation
PortOperation objects allows the user to alter the state of the port.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin configured state to disable the port.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
