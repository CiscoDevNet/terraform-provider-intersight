---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_mode"
description: |-
  Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
---

# Data Source: intersight_fabric_port_mode
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `custom_mode`:(string) Custom Port Mode specified for the port range.* `FibreChannel` - Fibre Channel Port Types.* `BreakoutEthernet10G` - Breakout Ethernet 10G Port Type.* `BreakoutEthernet25G` - Breakout Ethernet 25G Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id_end`:(int) Ending range of the Port Identifier. 
* `port_id_start`:(int) Starting range of the Port Identifier. 
* `slot_id`:(int) Slot Identifier of the switch. 
