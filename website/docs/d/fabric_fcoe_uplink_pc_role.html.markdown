
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_fcoe_uplink_pc_role"
sidebar_current: "docs-intersight-data-source-fabric-fcoe-uplink-pc-role"
description: |-
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
---

# Data Source: intersight_fabric._fcoe_uplink_pc_role
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `udld_admin_state`:(string) Admin configured state for UDLD for this port.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
