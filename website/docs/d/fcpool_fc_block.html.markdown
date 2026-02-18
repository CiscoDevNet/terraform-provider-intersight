---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_fc_block"
description: |-
        A Block object represents a sequential range of WWN (World Wide Name) addresses defined by a start and end address. A pool can contain multiple blocks.
        #### Purpose
        The block object serves as a fundamental unit within a WWN pool, enabling efficient management and allocation of contiguous network identifiers.
        #### Key Concepts
        - **Range Definition:** Represents a contiguous range of WWN addresses with defined start and end values.
        - **Identifier Type:** Manages WWN-based identifiers used in network fabrics.
        - **Pool Association:** Functions as a core component within WWN pools for structured resource management.
        - **Contiguity:** Ensures sequential and non-overlapping address allocation for optimal utilization.

---

# Data Source: intersight_fcpool_fc_block
A Block object represents a sequential range of WWN (World Wide Name) addresses defined by a start and end address. A pool can contain multiple blocks.
#### Purpose
The block object serves as a fundamental unit within a WWN pool, enabling efficient management and allocation of contiguous network identifiers.
#### Key Concepts
- **Range Definition:** Represents a contiguous range of WWN addresses with defined start and end values.
- **Identifier Type:** Manages WWN-based identifiers used in network fabrics.
- **Pool Association:** Functions as a core component within WWN pools for structured resource management.
- **Contiguity:** Ensures sequential and non-overlapping address allocation for optimal utilization.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_fc_block.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
