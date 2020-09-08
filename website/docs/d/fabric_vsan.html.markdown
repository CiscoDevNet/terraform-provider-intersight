
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `default_zoning`:(string) Enables or Disables the default zoning state.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `fc_zone_sharing_mode`:(string) Logical grouping mode for fc ports. 
* `fcoe_vlan`:(int) FCOE Vlan associated to the VSAN configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User given name for the VSAN configuration. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `vsan_id`:(int) Virtual San Identifier in the switch. 
