
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_operation"
sidebar_current: "docs-intersight-data-source-fabric-port-operation"
description: |-
PortOperation objects allows the user to alter the state of the port.
---

# Data Source: intersight_fabric._port_operation
PortOperation objects allows the user to alter the state of the port.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin configured state to disable the port.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
