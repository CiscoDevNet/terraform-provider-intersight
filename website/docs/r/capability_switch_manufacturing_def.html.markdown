---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_manufacturing_def"
description: |-
  Switch/Fabric-Interconnect manufacturing def properties.
---

# Resource: intersight_capability_switch_manufacturing_def
Switch/Fabric-Interconnect manufacturing def properties.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `caption`:(string) Caption for Switch/Fabric-Interconnect. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description for Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `part_number`:(string) Part Number for Switch/Fabric-Interconnect. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `product_name`:(string) Product Name for Switch/Fabric-Interconnect. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 


## Import
`intersight_capability_switch_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_switch_manufacturing_def.example 1234567890987654321abcde
```