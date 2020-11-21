
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_vsan"
sidebar_current: "docs-intersight-data-source-fabric-vsan"
description: |-
Configuration object sent by user to create VSAN configurations.
---

# Data Source: intersight_fabric._vsan
Configuration object sent by user to create VSAN configurations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `default_zoning`:(string) Enables or Disables the default zoning state.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `fc_zone_sharing_mode`:(string) Logical grouping mode for fc ports. 
* `fcoe_vlan`:(int) FCOE Vlan associated to the VSAN configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User given name for the VSAN configuration. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `vsan_id`:(int) Virtual San Identifier in the switch. 
