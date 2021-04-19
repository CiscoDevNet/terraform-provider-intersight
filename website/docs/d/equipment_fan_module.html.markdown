---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan_module"
description: |-
  This represents Fan module housing multiple fans for chassis/server.
---

# Data Source: intersight_equipment_fan_module
This represents Fan module housing multiple fans for chassis/server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_fan_module.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description for the fan module. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `module_id`:(int) This field acts as the identifier for this particular Module, within the Fabric Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) This field is used to indicate this fan module's operational state. 
* `part_number`:(string) This field identifies the Part Number for this Fan Module. 
* `pid`:(string) This field identifies the Product ID for the fan module. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) This field identifies the Stockkeeping Unit for this Fan Module. 
* `tray_id`:(int) Tray identifier for the fan module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for this Fan Module. 
 