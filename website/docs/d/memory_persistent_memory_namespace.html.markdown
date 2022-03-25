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
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_namespace.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(string) Capacity in GiB of the Persistent Memory Namespace. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health_state`:(string) Health state of the Persistent Memory Namespace. 
* `label_version`:(string) Label version of the Persistent Memory Namespace. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Mode of the Persistent Memory Namespace. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Persistent Memory Namespace. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the Persistent Memory Namespace. 
 
