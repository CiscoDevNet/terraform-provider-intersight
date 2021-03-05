---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_config_result"
description: |-
  Result of a previously applied Persistent Memory configuration on a server.
---

# Data Source: intersight_memory_persistent_memory_config_result
Result of a previously applied Persistent Memory configuration on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_config_result.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_error_desc`:(string) Error in the result of a previously applied Persistent Memory configuration on a server. 
* `config_result`:(string) Result of a previously applied Persistent Memory configuration on a server. 
* `config_sequence_no`:(int) Sequence number of a previously applied Persistent Memory configuration on a server. 
* `config_state`:(string) State of a previously applied Persistent Memory configuration on a server. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 