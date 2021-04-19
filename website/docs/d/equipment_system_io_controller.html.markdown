---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_system_io_controller"
description: |-
  I/O Controller on a chassis which provides the data path to S-series server.
---

# Data Source: intersight_equipment_system_io_controller
I/O Controller on a chassis which provides the data path to S-series server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_system_io_controller.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chassis_id`:(string) The assigned identifier for a chassis. 
* `connection_path`:(string) Connection Path identifies the data path available between IOModule and FI. 
* `connection_status`:(string) Connection status identifies the status of data path. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field gives a brief information on systemIOController. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `managing_instance`:(string) This field identifies the CIMC that manages the controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) This field identifies the SIOC operational state. 
* `part_number`:(string) Part Number identifier for the IO module. 
* `pid`:(string) This field identifies the Product ID for systemIOController. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `system_io_controller_id`:(int) This represents system I/O Controller identifier. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 