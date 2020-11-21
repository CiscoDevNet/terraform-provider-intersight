
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_operation"
sidebar_current: "docs-intersight-data-source-fabric-pc-operation"
description: |-
PcOperation objects allows the user to alter the state of the port channel.
---

# Data Source: intersight_fabric._pc_operation
PcOperation objects allows the user to alter the state of the port channel.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin configured state to disable the port channel.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
