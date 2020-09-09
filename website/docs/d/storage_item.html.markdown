
---
layout: "intersight"
page_title: "Intersight: intersight_storage_item"
sidebar_current: "docs-intersight-data-source-storage-item"
description: |-
FI Local Storage information.
---

# Data Source: intersight_storage._item
FI Local Storage information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `alarm_type`:(string) The alarmType of the Local storage in FI. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Local storage in FI. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) The operState of the Local storage in FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `size`:(string) The size (MB) of the Local storage in FI. 
* `used`:(string) The used percent of the Local storage in FI. 
