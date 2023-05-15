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
To access the ith object of the results obtained, use `data.intersight_equipment_psu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description for the power supply unit. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field is to provide name for the power supply unit. 
* `oper_state`:(string) This field identifies the psu operational state. 
* `part_number`:(string) This field identifies the Part Number for this Power Supply Unit. 
* `pid`:(string) This field identifies the Product ID for the Power Supply. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `psu_fw_version`:(string) This field identifies the Firmware Version of the Power Supply. 
* `psu_id`:(int) This represents power supply unit identifier in chassis/server. 
* `psu_input_src`:(string) This field identifies the input source for the Power Supply. 
* `psu_type`:(string) This field identifies the type of the Power Supply. 
* `psu_wattage`:(string) This field identifies the Wattage of the Power Supply. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Power Supply. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vid`:(string) This field identifies the Vendor ID for this Power Supply Unit. 
* `voltage`:(string) This field is used to indicate the voltage state for this Power Supply. 
 
