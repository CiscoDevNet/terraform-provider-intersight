
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_uplink_role"
sidebar_current: "docs-intersight-data-source-fabric-fc-uplink-role"
description: |-
Configuration object sent by user to create a fc uplink port.
---

# Data Source: intersight_fabric._fc_uplink_role
Configuration object sent by user to create a fc uplink port.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `8Gbps` - Admin configurable speed 8Gbps.* `16Gbps` - Admin configurable speed 16Gbps.* `32Gbps` - Admin configurable speed 32Gbps. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `fill_pattern`:(string) Fill pattern to differentiate the configs in NPIV.* `Idle` - Fc Fill Pattern type Idle.* `Arbff` - Fc Fill Pattern type Arbff. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
* `vsan_id`:(int) Virtual San Identifier associated to the FC port. 
