---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_psu"
description: |-
  This represents power supply unit for chassis/server.
---

# Data Source: intersight_equipment_psu
This represents power supply unit for chassis/server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_psu.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) This field is to provide description for the power supply unit. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
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
* `voltage`:(string) This field is used to indicate the voltage state for this Power Supply. 
 