---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_pool_member"
description: |-
        PoolMembers represent individual MAC addresses that are part of a specific pool. They provide per-address membership tracking and link to the corresponding “universe” lease record for correlation across pool and account-level bookkeeping.
        #### Purpose
        Offer a pool-scoped view of each MAC identity in the pool, including how it maps to the universe lease and which entity (if any) currently owns/uses the MAC.
        #### Key Concepts
        - **Per-pool identity membership**: Each PoolMember is uniquely identified by the combination of `pool` and `macAddress`.
        - **Universe correlation**: The `peer` relationship links the pool member to the corresponding `Lease` record in the universe.
        - **Block provenance**: `blockHead` indicates which IdBlock/range the MAC came from.
        - **Ownership linkage**: `assignedToEntity` ties the MAC to the consuming entity (e.g., server profile), enabling “who is using this MAC” queries.
        - **Pool relationship semantics**: `pool` is read-only and cascades on pool deletion, ensuring membership records are cleaned up with the pool.

---

# Data Source: intersight_macpool_pool_member
PoolMembers represent individual MAC addresses that are part of a specific pool. They provide per-address membership tracking and link to the corresponding “universe” lease record for correlation across pool and account-level bookkeeping.
#### Purpose
Offer a pool-scoped view of each MAC identity in the pool, including how it maps to the universe lease and which entity (if any) currently owns/uses the MAC.
#### Key Concepts
- **Per-pool identity membership**: Each PoolMember is uniquely identified by the combination of `pool` and `macAddress`.
- **Universe correlation**: The `peer` relationship links the pool member to the corresponding `Lease` record in the universe.
- **Block provenance**: `blockHead` indicates which IdBlock/range the MAC came from.
- **Ownership linkage**: `assignedToEntity` ties the MAC to the consuming entity (e.g., server profile), enabling “who is using this MAC” queries.
- **Pool relationship semantics**: `pool` is read-only and cascades on pool deletion, ensuring membership records are cleaned up with the pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_pool_member.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assigned`:(bool) Boolean to represent whether the ID is in use. 
* `assigned_by_another`:(bool) Boolean to represent whether the ID is used either statically or by another pool. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mac_address`:(string) MAC Address of this pool member. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reserved`:(bool) Boolean to represent whether the ID is reserved. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
