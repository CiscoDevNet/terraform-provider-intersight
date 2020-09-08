
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_transceiver"
sidebar_current: "docs-intersight-data-source-equipment-transceiver"
description: |-
Transceiver information on the chassis.
---

# Data Source: intersight_equipment._transceiver
Transceiver information on the chassis.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_speed`:(string) Operational speed of the transceiver. 
* `oper_state`:(string) Operational state of the transceiver. 
* `oper_state_qual`:(string) Reason for this transceiver's operational state. 
* `port_id`:(int) Switch physical port identifier. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `type`:(string) The type of the transceiver. 
* `vendor`:(string) This field identifies the vendor of the given component. 
