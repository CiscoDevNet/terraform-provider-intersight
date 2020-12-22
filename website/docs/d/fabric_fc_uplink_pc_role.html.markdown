---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_uplink_pc_role"
description: |-
  Object sent by user to configure a fc uplink port-channel on the collection of ports.
---

# Data Source: intersight_fabric_fc_uplink_pc_role
Object sent by user to configure a fc uplink port-channel on the collection of ports.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `8Gbps` - Admin configurable speed 8Gbps.* `16Gbps` - Admin configurable speed 16Gbps.* `32Gbps` - Admin configurable speed 32Gbps. 
* `fill_pattern`:(string) Fill pattern to differentiate the configs in NPIV.* `Idle` - Fc Fill Pattern type Idle.* `Arbff` - Fc Fill Pattern type Arbff. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `vsan_id`:(int) Virtual San Identifier associated to the FC port. 
