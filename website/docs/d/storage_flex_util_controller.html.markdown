---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_flex_util_controller"
description: |-
  Storage Flex Util Adapter.
---

# Data Source: intersight_storage_flex_util_controller
Storage Flex Util Adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_flex_util_controller.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `controller_name`:(string) Name of the Flex Util Controller. 
* `controller_status`:(string) The current status of the controller. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ff_controller_id`:(string) Identifier for the Storage Flex Util Controller. 
* `internal_state`:(string) The internal state of the controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 