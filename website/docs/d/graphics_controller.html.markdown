
---
layout: "intersight"
page_title: "Intersight: intersight_graphics_controller"
sidebar_current: "docs-intersight-data-source-graphics-controller"
description: |-
Controller for a Graphics Card.
---

# Data Source: intersight_graphics._controller
Controller for a Graphics Card.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `controller_id`:(int) The id of the graphics controller. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pci_addr`:(string) The PCI address of the graphics controller. 
* `pci_slot`:(string) The PCI slot information of the graphics controller. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
