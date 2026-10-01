---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_reservation"
description: |-
        Reservations hold MAC addresses that are set aside (reserved) so they won’t be allocated to other consumers. They can be created before a pool exists or before a particular identity is available, and later reconciled to the pool/block they belong to once that pool is created or discovered.
        #### Purpose
        Guarantee that specific MAC identities are preserved for intended consumers, supporting deterministic addressing, pre-provisioning workflows, and collision avoidance.
        #### Key Concepts
        - **Reserved identity**: `identity` is the MAC being reserved and is create-only to prevent “moving” a reservation to a different MAC.
        - **Deferred pool association**: `memberOf` (read-only) can be populated at reservation creation (if the pool exists) or later during pool creation, linking to the pool and block Moids.
        - **Pool/block/member relationships**:
        - `pool` links the reservation to a Pool (read-write reference).
        - `blockHead` links to the IdBlock containing the reserved MAC (unset on peer delete).
        - `poolMember` links to the corresponding PoolMember (unset on peer delete).
        - **Universe context**: `universe` references the MAC universe bookkeeping container.
        - **Access governance**: Supports create/delete privileges explicitly for MAC reservation operations, aligning with reservation-specific administrative controls.

---

# Data Source: intersight_macpool_reservation
Reservations hold MAC addresses that are set aside (reserved) so they won’t be allocated to other consumers. They can be created before a pool exists or before a particular identity is available, and later reconciled to the pool/block they belong to once that pool is created or discovered.
#### Purpose
Guarantee that specific MAC identities are preserved for intended consumers, supporting deterministic addressing, pre-provisioning workflows, and collision avoidance.
#### Key Concepts
- **Reserved identity**: `identity` is the MAC being reserved and is create-only to prevent “moving” a reservation to a different MAC.
- **Deferred pool association**: `memberOf` (read-only) can be populated at reservation creation (if the pool exists) or later during pool creation, linking to the pool and block Moids.
- **Pool/block/member relationships**:
  - `pool` links the reservation to a Pool (read-write reference).
  - `blockHead` links to the IdBlock containing the reserved MAC (unset on peer delete).
  - `poolMember` links to the corresponding PoolMember (unset on peer delete).
- **Universe context**: `universe` references the MAC universe bookkeeping container.
- **Access governance**: Supports create/delete privileges explicitly for MAC reservation operations, aligning with reservation-specific administrative controls.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_reservation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the allocation for the identity in the reservation either static or dynamic (i.e. via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) MAC identity to be reserved. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
