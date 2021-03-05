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
* `chassis_id`:(string) The assigned identifier for a chassis. 
* `connection_path`:(string) Connection Path identifies the data path available between IOModule and FI. 
* `connection_status`:(string) Connection status identifies the status of data path. 
* `description`:(string) This field gives a brief information on systemIOController. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `managing_instance`:(string) This field identifies the CIMC that manages the controller. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) This field identifies the SIOC operational state. 
* `part_number`:(string) Part Number identifier for the IO module. 
* `pid`:(string) This field identifies the Product ID for systemIOController. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `system_io_controller_id`:(int) This represents system I/O Controller identifier. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 