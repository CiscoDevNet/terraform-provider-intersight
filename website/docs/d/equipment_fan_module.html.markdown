
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan_module"
sidebar_current: "docs-intersight-data-source-equipment-fan-module"
description: |-
This represents Fan module housing multiple fans for chassis/server.
---

# Data Source: intersight_equipment._fan_module
This represents Fan module housing multiple fans for chassis/server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) This field is to provide description for the fan module. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `module_id`:(int) This field acts as the identifier for this particular Module, within the Fabric Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) This field is used to indicate this fan module's operational state. 
* `part_number`:(string) This field identifies the Part Number for this Fan Module. 
* `pid`:(string) This field identifies the Product ID for the fan module. 
* `presence`:(string) This field is used to indicate this fan module's presence. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Fan Module. 
* `tray_id`:(int) Tray identifier for the fan module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for this Fan Module. 
