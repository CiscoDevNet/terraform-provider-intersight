
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_psu"
sidebar_current: "docs-intersight-data-source-equipment-psu"
description: |-
This represents power supply unit for chassis/server.
---

# Data Source: intersight_equipment._psu
This represents power supply unit for chassis/server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) This field is to provide description for the power supply unit. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) This field identifies the psu operational state. 
* `part_number`:(string) This field identifies the Part Number for this Power Supply Unit. 
* `pid`:(string) This field identifies the Product ID for the Power Supply. 
* `presence`:(string) This field identifies the presence state of the psu. 
* `psu_fw_version`:(string) This field identifies the Firmware Version of the Power Supply. 
* `psu_id`:(int) This represents power supply unit identifier in chassis/server. 
* `psu_input_src`:(string) This field identifies the input source for the Power Supply. 
* `psu_type`:(string) This field identifies the type of the Power Supply. 
* `psu_wattage`:(string) This field identifies the Wattage of the Power Supply. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Power Supply. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for this Power Supply Unit. 
* `voltage`:(string) This field is used to indicate the Voltage for this Power Supply. 
