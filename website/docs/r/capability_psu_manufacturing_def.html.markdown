---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_psu_manufacturing_def"
description: |-
  Power supply unit properties.
---

# Resource: intersight_capability_psu_manufacturing_def
Power supply unit properties.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `caption`:(string) Caption for a power supply unit. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description for a power supply unit. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pid`:(string) Product Identifier for a power supply unit. 
* `product_name`:(string) Product Name for Power Supplu Unit. 
* `sku`:(string) SKU information for a power supply unit. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for a power supply unit. 


## Import
`intersight_capability_psu_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_psu_manufacturing_def.example 1234567890987654321abcde
```