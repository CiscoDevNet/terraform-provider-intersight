
---
layout: "intersight"
page_title: "Intersight: intersight_inventory_generic_inventory_holder"
sidebar_current: "docs-intersight-data-source-inventory-generic-inventory-holder"
description: |-
A container class for generic inventory.
---

# Data Source: intersight_inventory._generic_inventory_holder
A container class for generic inventory.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `endpoint`:(string) The endpoint represented by this holder. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
