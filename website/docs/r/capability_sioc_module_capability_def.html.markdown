---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_sioc_module_capability_def"
description: |-
  Chassis SIOC module capabilities.
---

# Resource: intersight_capability_sioc_module_capability_def
Chassis SIOC module capabilities.
## Argument Reference
The following arguments are supported:
* `dc_supported`:(bool) Device connector support on SIOC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_capability_sioc_module_capability_def` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_sioc_module_capability_def.example 1234567890987654321abcde
``` 