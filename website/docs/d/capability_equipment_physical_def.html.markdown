
---
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_physical_def"
sidebar_current: "docs-intersight-data-source-capability-equipment-physical-def"
description: |-
Type to represent additional switch specific capabilities.
---

# Data Source: intersight_capability._equipment_physical_def
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `depth`:(float) Depth information for a Switch/Fabric-Interconnect. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect. 
* `max_power`:(int) Max Power information for a Switch/Fabric-Interconnect. 
* `min_power`:(int) Min Power information for a Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `nominal_power`:(int) Nominal Power information for a Switch/Fabric-Interconnect. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `weight`:(float) Weight information for a Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect. 
