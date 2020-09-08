
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_locator_led"
sidebar_current: "docs-intersight-data-source-equipment-locator-led"
description: |-
Locator Led of an Equipment like Rack, Disk etc.
---

# Data Source: intersight_equipment._locator_led
Locator Led of an Equipment like Rack, Disk etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `color`:(string) Color of the locatorled available on an equipment. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Identifies the operational state of locatorled. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
