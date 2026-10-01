---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_namespace_config_result"
description: |-
        PersistentMemoryNamespaces represent persistent memory namespaces created within a persistent memory region, including name, UUID, capacity, mode, and health.
        #### Purpose
        Provide inventory for PMem namespaces so users can see the logical PMem devices configured on the platform.
        #### Key Concepts
        - **Logical PMem device:** A namespace is a logical allocation within a region.
        - **Capacity/mode visibility:** Exposes size and operating mode.
        - **Identity:** Uses UUID and name for referencing and correlation.

---

# Data Source: intersight_memory_persistent_memory_namespace_config_result
PersistentMemoryNamespaces represent persistent memory namespaces created within a persistent memory region, including name, UUID, capacity, mode, and health.
#### Purpose
Provide inventory for PMem namespaces so users can see the logical PMem devices configured on the platform.
#### Key Concepts
- **Logical PMem device:** A namespace is a logical allocation within a region.
- **Capacity/mode visibility:** Exposes size and operating mode.
- **Identity:** Uses UUID and name for referencing and correlation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_namespace_config_result.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_status`:(string) Status of the Persistent Memory Namespace needed to be configured. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of a Persistent Memory Namespace that needed to be configured. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `socket_id`:(string) Socket ID in which the Persistent Memory Namespace needed to be configured. 
* `socket_memory_id`:(string) Socket Memory ID in which the Persistent Memory Namespace needed to be configured. 
 
