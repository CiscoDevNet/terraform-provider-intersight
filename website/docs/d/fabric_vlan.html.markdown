
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `is_native`:(bool) Used to define whether this VLAN is to be classified as 'native' for traffic in this FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The 'name' used to identify this VLAN. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `vlan_id`:(int) The identifier for this Virtual LAN. 
