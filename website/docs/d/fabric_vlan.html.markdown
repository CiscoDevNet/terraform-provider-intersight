
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_vlan"
sidebar_current: "docs-intersight-data-source-fabric-vlan"
description: |-
Configuration object for Virtual LAN.
---

# Data Source: intersight_fabric._vlan
Configuration object for Virtual LAN.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `is_native`:(bool) Used to define whether this VLAN is to be classified as 'native' for traffic in this FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The 'name' used to identify this VLAN. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `vlan_id`:(int) The identifier for this Virtual LAN. 
