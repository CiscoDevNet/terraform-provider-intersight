---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_block"
description: |-
        The Blocks object represents a contiguous range of identifiers that are part of a larger pool.
        #### Purpose
        It defines a specific range of addresses or identifiers within a pool, facilitating organized management and efficient allocation of resources.
        #### Key Concepts
        - **Range Definition:** Represents a contiguous range of identifiers with defined start and end values.
        - **Pool Association:** Functions as a core component within pools for structured resource management.
        - **Contiguity:** Ensures sequential and non-overlapping address allocation.

---

# Data Source: intersight_iqnpool_block
The Blocks object represents a contiguous range of identifiers that are part of a larger pool.
#### Purpose
It defines a specific range of addresses or identifiers within a pool, facilitating organized management and efficient allocation of resources.
#### Key Concepts
- **Range Definition:** Represents a contiguous range of identifiers with defined start and end values.
- **Pool Association:** Functions as a core component within pools for structured resource management.
- **Contiguity:** Ensures sequential and non-overlapping address allocation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iqnpool_block.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `prefix`:(string) Prefix of the IQN pool. IQN Address is constructed as <prefix>:<suffix>:<number>. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
