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
* `caption`:(string) Caption for a power supply unit. 
* `description`:(string) Description for a power supply unit. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a power supply unit. 
* `product_name`:(string) Product Name for Power Supplu Unit. 
* `sku`:(string) SKU information for a power supply unit. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for a power supply unit. 


## Import
`intersight_capability_psu_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_psu_manufacturing_def.example 1234567890987654321abcde
``` 