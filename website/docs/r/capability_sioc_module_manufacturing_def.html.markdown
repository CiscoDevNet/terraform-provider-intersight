---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_sioc_module_manufacturing_def"
description: |-
  Chassis SIOC module properties.
---

# Resource: intersight_capability_sioc_module_manufacturing_def
Chassis SIOC module properties.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `caption`:(string) Caption for a chassis SIOC module. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description for a chassis SIOC module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pid`:(string) Product Identifier for a chassis SIOC module. 
* `product_name`:(string) Product Name for SIOC Module. 
* `sku`:(string) SKU information for a chassis SIOC module. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for a chassis SIOC module. 


## Import
`intersight_capability_sioc_module_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_sioc_module_manufacturing_def.example 1234567890987654321abcde
```