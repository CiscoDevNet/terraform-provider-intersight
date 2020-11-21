
---
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_region"
sidebar_current: "docs-intersight-data-source-memory-persistent-memory-region"
description: |-
Persistent Memory Region configured on the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory._persistent_memory_region
Persistent Memory Region configured on the Persistent Memory Modules on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `free_capacity`:(string) Free capacity in GiB of the Persistent Memory Region. 
* `health_state`:(string) Health state of the Persistent Memory Region. 
* `interleaved_set_id`:(string) ID of the Interleaved Set formed for this Persistent Memory Region. 
* `locater_ids`:(string) Set of locator IDs that are included in the Persistent Memory Region. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `persistent_memory_type`:(string) Persistent Memory type of the Persistent Memory Region. 
* `region_id`:(string) ID of the Persistent Memory Region. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `socket_id`:(string) Socket ID of the Persistent Memory Region. 
* `socket_memory_id`:(string) Socket Memory ID of the Persistent Memory Region. 
* `total_capacity`:(string) Total capacity in GiB of the Persistent Memory Region. 
