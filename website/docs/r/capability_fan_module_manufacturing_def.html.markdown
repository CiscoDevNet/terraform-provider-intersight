---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_fan_module_manufacturing_def"
description: |-
  Fan module unit that contains multiple fans.
---

# Resource: intersight_capability_fan_module_manufacturing_def
Fan module unit that contains multiple fans.
## Argument Reference
The following arguments are supported:
* `caption`:(string) Caption for a fan module. 
* `description`:(string) Description for a fan module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a fan module. 
* `product_name`:(string) Product Name for Fan Module Unit. 
* `sku`:(string) SKU information for a fan module. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for a fan module. 


## Import
`intersight_capability_fan_module_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_fan_module_manufacturing_def.example 1234567890987654321abcde
``` 