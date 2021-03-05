---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_member"
description: |-
  PcMember object is to establish the relationship between port parameters and pcId.
---

# Data Source: intersight_fabric_pc_member
PcMember object is to establish the relationship between port parameters and pcId.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_pc_member.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
 