---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_item"
description: |-
  FI Local Storage information.
---

# Data Source: intersight_storage_item
FI Local Storage information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_item.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `alarm_type`:(string) The alarmType of the Local storage in FI. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Local storage in FI. 
* `oper_state`:(string) The operState of the Local storage in FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `size`:(string) The size (MB) of the Local storage in FI. 
* `used`:(string) The used percent of the Local storage in FI. 
 