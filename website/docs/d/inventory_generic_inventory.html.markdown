
---
layout: "intersight"
page_title: "Intersight: intersight_inventory_generic_inventory"
sidebar_current: "docs-intersight-data-source-inventory-generic-inventory"
description: |-
Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.
---

# Data Source: intersight_inventory._generic_inventory
Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `key`:(string) Key of inventory data for Generic Inventory data set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `type`:(string) Type of inventory data for Generic Inventory data set. 
* `value`:(string) Value of inventory data for Generic Inventory data set. 
