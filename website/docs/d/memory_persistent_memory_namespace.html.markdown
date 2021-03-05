---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_namespace"
description: |-
  Persistent Memory Namespace configured within a Persistent Memory Region on a server.
---

# Data Source: intersight_memory_persistent_memory_namespace
Persistent Memory Namespace configured within a Persistent Memory Region on a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_namespace.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(string) Capacity in GiB of the Persistent Memory Namespace. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `health_state`:(string) Health state of the Persistent Memory Namespace. 
* `label_version`:(string) Label version of the Persistent Memory Namespace. 
* `mode`:(string) Mode of the Persistent Memory Namespace. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Persistent Memory Namespace. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `uuid`:(string) UUID of the Persistent Memory Namespace. 
 