
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Local storage in FI. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `oper_state`:(string) The operState of the Local storage in FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `size`:(string) The size (MB) of the Local storage in FI. 
* `used`:(string) The used percent of the Local storage in FI. 
