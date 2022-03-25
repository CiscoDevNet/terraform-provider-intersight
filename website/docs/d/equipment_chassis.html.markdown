---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_chassis"
description: |-
        A physical holder housing blade servers.

---

# Data Source: intersight_equipment_chassis
A physical holder housing blade servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_chassis.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chassis_id`:(int) The assigned identifier for a chassis. 
* `connection_path`:(string) This field identifies the connectivity path for the chassis enclosure. 
* `connection_status`:(string) This field identifies the connectivity status for the chassis enclosure. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description for chassis model. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fault_summary`:(int) This field summarizes the faults on the chassis enclosure. 
* `management_mode`:(string) The management mode of the blade server chassis.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field identifies the name for the chassis enclosure. 
* `oper_state`:(string) This field identifies the Chassis Operational State. 
* `part_number`:(string) Part Number identifier for the chassis enclosure. 
* `pid`:(string) This field identifies the Product ID for the chassis enclosure. 
* `platform_type`:(string) The platform type that the chassis is a part of. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `product_name`:(string) This field identifies the Product Name for the chassis enclosure. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) This field identifies the Stock Keeping Unit for the chassis enclosure. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for the chassis enclosure. 
 
