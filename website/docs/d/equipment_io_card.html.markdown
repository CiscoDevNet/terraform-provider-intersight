---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_io_card"
description: |-
        I/O module on a chassis which multiplexes traffic from blade servers.

---

# Data Source: intersight_equipment_io_card
I/O module on a chassis which multiplexes traffic from blade servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_io_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connection_path`:(string) Switch Id to which the IOM is connected to. The value can be A or B. 
* `connection_status`:(string) Connectivity Status of FEX/IOM to Switch - A or B or AB. 
* `create_time`:(string) The time when this managed object was created. 
* `dc_supported`:(bool) IOM device connector support. 
* `description`:(string) This field is to provide description for the iocard module model. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `module_id`:(int) Module Identifier for the IO module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of IO card or fabric extender. 
* `part_number`:(string) Part Number identifier for the IO module. 
* `pid`:(string) This field identifies the Product ID for the IO module. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `product_name`:(string) This field identifies the Product Name for the iocard module model. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `side`:(string) Location of IOM within a chassis. The value can be left or right. 
* `sku`:(string) This field identifies the Stock Keeping Unit for the IO card module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) This field identifies the version of the IO card module. 
* `vid`:(string) This field identifies the Vendor ID for the IO card module. 
 
