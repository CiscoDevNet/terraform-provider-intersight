---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_port_group_aggregation_def"
description: |-
  FEX/IOCARD module port group aggregation capabilities.
---

# Resource: intersight_capability_port_group_aggregation_def
FEX/IOCARD module port group aggregation capabilities.
## Argument Reference
The following arguments are supported:
* `aggregation_cap`:(string) Aggregation capability for port group. 
* `hw40_g_port_group_cap`:(bool) Indicates support for 40G port group capability. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pgtype`:(string) The type of port group for which this capability is defined. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_capability_port_group_aggregation_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_port_group_aggregation_def.example 1234567890987654321abcde
``` 