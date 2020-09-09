
---
layout: "intersight"
page_title: "Intersight: intersight_graphics_card"
sidebar_current: "docs-intersight-data-source-graphics-card"
description: |-
Graphics Card present in a server.
---

# Data Source: intersight_graphics._card
Graphics Card present in a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `card_id`:(int) The id of the graphics card. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_id`:(int) The device id of the graphics card. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `expander_slot`:(string) The expander slot information of the card. 
* `firmware_version`:(string) The firmware version of the graphics card. 
* `mode`:(string) The current mode of the graphics card. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_gpus`:(string) The number of controllers under each card. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) The current operational state of the graphics card. 
* `pci_address`:(string) The PCI address of the graphics card. 
* `pci_address_list`:(string) This list contains the PCI address of all controllers for corresponding card. 
* `pci_slot`:(string) The PCI slot name of the graphics card. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
