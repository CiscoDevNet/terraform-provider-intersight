---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_configuration"
description: |-
  Persistent Memory configuration on all the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory_persistent_memory_configuration
Persistent Memory configuration on all the Persistent Memory Modules on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_configuration.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `memory_capacity`:(string) Memory capacity in GiB of a Persistent Memory configuration on a server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_modules`:(string) Number of Persistent Memory Modules on a server. 
* `num_of_regions`:(string) Number of Persistent Memory Regions on a server. 
* `persistent_memory_capacity`:(string) Persistent memory capacity in GiB of a Persistent Memory configuration on a server. 
* `reserved_capacity`:(string) Reserved capacity in GiB of a Persistent Memory configuration on a server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `security_state`:(string) Collective security state of all Persistent Memory modules on a server. 
* `total_capacity`:(string) Total capacity in GiB of a Persistent Memory configuration on a server. 
 