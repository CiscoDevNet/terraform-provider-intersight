
---
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_configuration"
sidebar_current: "docs-intersight-data-source-memory-persistent-memory-configuration"
description: |-
Persistent Memory configuration on all the Persistent Memory Modules on a server.
---

# Data Source: intersight_memory._persistent_memory_configuration
Persistent Memory configuration on all the Persistent Memory Modules on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `memory_capacity`:(string) Memory capacity in GiB of a Persistent Memory configuration on a server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_modules`:(string) Number of Persistent Memory Modules on a server. 
* `num_of_regions`:(string) Number of Persistent Memory Regions on a server. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `persistent_memory_capacity`:(string) Persistent memory capacity in GiB of a Persistent Memory configuration on a server. 
* `reserved_capacity`:(string) Reserved capacity in GiB of a Persistent Memory configuration on a server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `security_state`:(string) Collective security state of all Persistent Memory modules on a server. 
* `total_capacity`:(string) Total capacity in GiB of a Persistent Memory configuration on a server. 
