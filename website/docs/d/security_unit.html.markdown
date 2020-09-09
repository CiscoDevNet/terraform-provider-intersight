
---
layout: "intersight"
page_title: "Intersight: intersight_security_unit"
sidebar_current: "docs-intersight-data-source-security-unit"
description: |-
The crypto card present on a server.
---

# Data Source: intersight_security._unit
The crypto card present on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Operational state of the security unit. 
* `operability`:(string) Operability state of the security unit. 
* `part_number`:(string) The part number of the security unit. 
* `pci_slot`:(string) PCIe slot of the security unit in the server. 
* `power`:(string) Power state of the security unit. 
* `presence`:(string) Security unit presence (equipped) or absence. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `thermal`:(string) Thermal state of the security unit. 
* `unit_id`:(int) The unique identifier assigned to the security unit within the server. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) The vendor identifier of the security unit. 
* `voltage`:(string) The voltage state of the security unit. 
