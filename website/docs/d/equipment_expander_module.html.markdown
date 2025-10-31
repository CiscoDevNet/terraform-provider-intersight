---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_expander_module"
description: |-
        Expander module inside the chassis.

---

# Data Source: intersight_equipment_expander_module
Expander module inside the chassis.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_expander_module.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) Firmware Version of the Chassis expander module. 
* `inventory_ready`:(bool) The inventory ready field indicates whether the chassis expander module management controller has completed inventory of the installed devices. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `module_id`:(int) Module identifier for the expander module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of expander module. 
* `part_number`:(string) Part number identifier for the expander module. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `side`:(string) Location of the expander module within a chassis. The value can be top or bottom.* `unknown` - Physical location of the module is unknown.* `top` - Physical location of the module is on the top part of the chassis.* `bottom` - Physical location of the module is on the bottom part of the chassis.* `left` - Physical location of the module is on the left side of the chassis.* `right` - Physical location of the module is on the right side of the chassis. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
