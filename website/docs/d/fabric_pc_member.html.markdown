
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_member"
sidebar_current: "docs-intersight-data-source-fabric-pc-member"
description: |-
PcMember object is to establish the relationship between port parameters and pcId.
---

# Data Source: intersight_fabric._pc_member
PcMember object is to establish the relationship between port parameters and pcId.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
