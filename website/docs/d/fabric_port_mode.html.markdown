
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_mode"
sidebar_current: "docs-intersight-data-source-fabric-port-mode"
description: |-
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
---

# Data Source: intersight_fabric._port_mode
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `custom_mode`:(string) Custom Port Mode specified for the port range.* `FibreChannel` - Fibre Channel Port Types.* `BreakoutEthernet10G` - Breakout Ethernet 10G Port Type.* `BreakoutEthernet25G` - Breakout Ethernet 25G Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `port_id_end`:(int) Ending range of the Port Identifier. 
* `port_id_start`:(int) Starting range of the Port Identifier. 
* `slot_id`:(int) Slot Identifier of the switch. 
