---
subcategory: "management"
layout: "intersight"
page_title: "Intersight: intersight_management_controller"
description: |-
  A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.
---

# Data Source: intersight_management_controller
A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_management_controller.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) Model of the endpoint that houses the management controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 