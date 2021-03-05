---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_region"
description: |-
  Persistent Memory Region configured on the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory_persistent_memory_region
Persistent Memory Region configured on the Persistent Memory Modules on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_region.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `free_capacity`:(string) Free capacity in GiB of the Persistent Memory Region. 
* `health_state`:(string) Health state of the Persistent Memory Region. 
* `interleaved_set_id`:(string) ID of the Interleaved Set formed for this Persistent Memory Region. 
* `locater_ids`:(string) Set of locator IDs that are included in the Persistent Memory Region. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `persistent_memory_type`:(string) Persistent Memory type of the Persistent Memory Region. 
* `region_id`:(string) ID of the Persistent Memory Region. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `socket_id`:(string) Socket ID of the Persistent Memory Region. 
* `socket_memory_id`:(string) Socket Memory ID of the Persistent Memory Region. 
* `total_capacity`:(string) Total capacity in GiB of the Persistent Memory Region. 
 