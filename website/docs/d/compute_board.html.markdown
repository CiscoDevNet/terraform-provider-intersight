---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_board"
description: |-
  Mother board of a server.
---

# Data Source: intersight_compute_board
Mother board of a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_board.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `board_id`:(int) Unique identifier of the mother board present in the server. 
* `cpu_type_controller`:(string) The type of central processing unit on the mother board. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_power_state`:(string) Current power state of the mother board of the server. 
* `presence`:(string) Identifies the presence of the mother board of the server. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 