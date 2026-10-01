---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_pool"
description: |-
        Pools represent a collection of MAC addresses that can be allocated to consumers (for example, VNICs associated with server profiles). A pool defines one or more MAC blocks that collectively make up the available allocation space.
        #### Purpose
        Provide the administrative construct for defining, managing, and reusing MAC address ranges across policies/profiles, enabling consistent and conflict-free MAC assignment at scale.
        #### Key Concepts
        - **Administrative container for MAC ranges**: `macBlocks` defines one or more address ranges (blocks) that compose the pool.
        - **System-managed block objects**: `blockHeads` provides read-only pointers to the realized IdBlock objects created for the pool.
        - **Lifecycle controls and tagging**: Supports CRUD operations and conditional privileges for setting/unsetting tags (policy governance).
        - **Stable identity**: The pool is identified by `name`, allowing consistent references from reservations and allocation flows.
        - **Downstream consumers**: The pool exists to serve allocations (leases) and to back pool membership tracking (pool members).

---

# Data Source: intersight_macpool_pool
Pools represent a collection of MAC addresses that can be allocated to consumers (for example, VNICs associated with server profiles). A pool defines one or more MAC blocks that collectively make up the available allocation space.
#### Purpose
Provide the administrative construct for defining, managing, and reusing MAC address ranges across policies/profiles, enabling consistent and conflict-free MAC assignment at scale.
#### Key Concepts
- **Administrative container for MAC ranges**: `macBlocks` defines one or more address ranges (blocks) that compose the pool.
- **System-managed block objects**: `blockHeads` provides read-only pointers to the realized IdBlock objects created for the pool.
- **Lifecycle controls and tagging**: Supports CRUD operations and conditional privileges for setting/unsetting tags (policy governance).
- **Stable identity**: The pool is identified by `name`, allowing consistent references from reservations and allocation flows.
- **Downstream consumers**: The pool exists to serve allocations (leases) and to back pool membership tracking (pool members).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assigned`:(int) Number of IDs that are currently assigned (in use). 
* `assignment_order`:(string) Property is deprecated. Sequential is the only assignment order supported.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `reserved`:(int) Number of IDs that are currently reserved (and not in use). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Total number of identifiers in this pool. 
 
