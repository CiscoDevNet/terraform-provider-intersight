
---
layout: "intersight"
page_title: "Intersight: intersight_port_mac_binding"
sidebar_current: "docs-intersight-data-source-port-mac-binding"
description: |-
Establishes relationship between the ports and connected end points based on LLDP TLVs.
---

# Data Source: intersight_port._mac_binding
Establishes relationship between the ports and connected end points based on LLDP TLVs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aggregate_port_id`:(int) Aggregate Port ID of the local Switch Interface. 
* `chassis_id`:(int) Chassis/FEX device idetifier that is local to a cluster. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mac`:(string) Device ID value that is advertised and available as a part of LLDP TLV. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `port_id`:(int) Port ID of the local Switch Interface. 
* `port_mac`:(string) Port ID value that is advertised and available as a part of LLDP TLV. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot_id`:(int) Slot ID of the local Switch slot Interface. 
* `switch_id`:(int) Switch Identifier that is local to a cluster. 
