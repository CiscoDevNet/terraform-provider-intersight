
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan"
sidebar_current: "docs-intersight-data-source-equipment-fan"
description: |-
Fan in a Fabric Interconnect / Chassis / RackUnit.
---

# Data Source: intersight_equipment._fan
Fan in a Fabric Interconnect / Chassis / RackUnit.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) This field is to provide description for the fan. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `fan_id`:(int) This field acts as the identifier for this particular Fan, within the Fabric Interconnect. 
* `fan_module_id`:(int) This field is used to identify the Fan Module to which this Fan belongs. 
* `model`:(string) This field identifies the model of the given component. 
* `module_id`:(int) Fan module Identifier for the fan. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `oper_state`:(string) This field is used to indicate this fan unit's operational state. 
* `part_number`:(string) This field identifies the Part Number for this Fan Unit. 
* `pid`:(string) This field identifies the Product ID for the fans. 
* `presence`:(string) This field is used to indicate this fan unit's presence. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Fan Unit. 
* `tray_id`:(int) Tray identifier for the fan module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for this Fan Unit. 
