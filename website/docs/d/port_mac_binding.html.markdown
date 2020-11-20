
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mac`:(string) Device ID value that is advertised and available as a part of LLDP TLV. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `port_id`:(int) Port ID of the local Switch Interface. 
* `port_mac`:(string) Port ID value that is advertised and available as a part of LLDP TLV. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot_id`:(int) Slot ID of the local Switch slot Interface. 
* `switch_id`:(int) Switch Identifier that is local to a cluster. 
