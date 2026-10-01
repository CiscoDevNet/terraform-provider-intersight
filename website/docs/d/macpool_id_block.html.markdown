---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_id_block"
description: |-
        IdBlocks represent contiguous blocks of MAC addresses that belong to a specific MAC pool. They act as the pool’s concrete address supply, defining the exact from/to range that can be allocated, reserved, and tracked.
        #### Purpose
        Provide a manageable, queryable block construct for MAC pools so the system can allocate MACs efficiently, validate uniqueness, and correlate pool membership and reservations back to a specific contiguous range.
        #### Key Concepts
        - **Contiguous range definition**: Each IdBlock encapsulates a `macBlock` (with `from` and `to`) describing the MAC range.
        - **Pool-scoped supply**: Blocks are correlated to a pool (index on `pool.Moid`) and are the source of addresses for leases and pool members.
        - **Uniqueness and lookup**: Indexing on `macBlock.From`/`macBlock.To` supports fast range queries and uniqueness checks.
        - **Permission inheritance**: Inherits permissions from the associated `pool`, aligning access control with pool ownership.

---

# Data Source: intersight_macpool_id_block
IdBlocks represent contiguous blocks of MAC addresses that belong to a specific MAC pool. They act as the pool’s concrete address supply, defining the exact from/to range that can be allocated, reserved, and tracked.
#### Purpose
Provide a manageable, queryable block construct for MAC pools so the system can allocate MACs efficiently, validate uniqueness, and correlate pool membership and reservations back to a specific contiguous range.
#### Key Concepts
- **Contiguous range definition**: Each IdBlock encapsulates a `macBlock` (with `from` and `to`) describing the MAC range.
- **Pool-scoped supply**: Blocks are correlated to a pool (index on `pool.Moid`) and are the source of addresses for leases and pool members.
- **Uniqueness and lookup**: Indexing on `macBlock.From`/`macBlock.To` supports fast range queries and uniqueness checks.
- **Permission inheritance**: Inherits permissions from the associated `pool`, aligning access control with pool ownership.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_id_block.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
