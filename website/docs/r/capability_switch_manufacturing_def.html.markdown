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
* `caption`:(string) Caption for Switch/Fabric-Interconnect. 
* `description`:(string) Description for Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `part_number`:(string) Part Number for Switch/Fabric-Interconnect. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `product_name`:(string) Product Name for Switch/Fabric-Interconnect. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 


## Import
`intersight_capability_switch_manufacturing_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_switch_manufacturing_def.example 1234567890987654321abcde
``` 