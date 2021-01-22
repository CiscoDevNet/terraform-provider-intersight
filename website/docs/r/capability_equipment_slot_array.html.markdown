---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_slot_array"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability_equipment_slot_array
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `first_index`:(float) First Index information for a Switch/Fabric-Interconnect hardware. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect hardware. 
* `horizontal_start_offset`:(float) Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_separation`:(float) Inline Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_size`:(float) Inline Group Size information for a Switch/Fabric-Interconnect hardware. 
* `inline_offset`:(float) Inline Offset information for a Switch/Fabric-Interconnect hardware. 
* `location`:(string) Location information for a Switch/Fabric-Interconnect hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `number_of_slots`:(int) Number of Slots information for a Switch/Fabric-Interconnect hardware. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `orientation`:(string) Orientation information for a Switch/Fabric-Interconnect hardware. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `selector`:(string) Selector information for a Switch/Fabric-Interconnect hardware. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `slots_per_line`:(int) Slots per line information for a Switch/Fabric-Interconnect hardware. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `transverse_group_separation`:(float) Transverse Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `transverse_group_size`:(float) Transverse Group Size information for a Switch/Fabric-Interconnect hardware. 
* `transverse_offset`:(float) Transverse Offset information for a Switch/Fabric-Interconnect hardware. 
* `vertical_start_offset`:(float) Vertical Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect hardware. 


## Import
`intersight_capability_equipment_slot_array` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_equipment_slot_array.example 1234567890987654321abcde
```