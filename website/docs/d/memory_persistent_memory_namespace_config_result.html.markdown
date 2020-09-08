
---
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_namespace_config_result"
sidebar_current: "docs-intersight-data-source-memory-persistent-memory-namespace-config-result"
description: |-
Result of a previously configured Persistent Memory Namespace on a server.
---

# Data Source: intersight_memory._persistent_memory_namespace_config_result
Result of a previously configured Persistent Memory Namespace on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_status`:(string) Status of the Persistent Memory Namespace needed to be configured. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of a Persistent Memory Namespace that needed to be configured. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `socket_id`:(string) Socket ID in which the Persistent Memory Namespace needed to be configured. 
* `socket_memory_id`:(string) Socket Memory ID in which the Persistent Memory Namespace needed to be configured. 
