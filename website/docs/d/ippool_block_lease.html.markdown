---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_block_lease"
description: |-
        BlockLeases represent an IP address allocation record issued from an IP pool to a specific consuming entity (for example, a server profile). The object acts as the “holder” for an allocated address (or block context) and provides the correlation points needed to track which universe/VRF scope the address belongs to and which entity currently owns it.
        #### Purpose
        Track and manage IP allocations at the block level so the system can enforce uniqueness, associate the allocation with the correct IP universe and VRF scope, and relate the allocation to the set of underlying IP lease objects.
        #### Key Concepts
        - **Pool-based allocation record**: Represents an IP address allocated from a pool for use by a specific consumer.
        - **Universe-scoped bookkeeping**: The `universe` relationship anchors the allocation in the IP Universe (with cascade on universe deletion) and is used for permission inheritance.
        - **VRF-scoped allocation**: `vrf` is `createonly`, ensuring the routing context for the allocation is fixed at creation time.
        - **Ownership correlation**: `assignedToEntity` links the allocation to the consuming managed object (e.g., server profile) for traceability and lifecycle management.
        - **Lease grouping**: `ipLeases` provides a collection of related `IpLease` objects, enabling a single block-level record to reference one or more concrete lease entries.
        - **System-managed lifecycle**: While users can READ, CREATE/UPDATE/DELETE are system API methods, reflecting that allocations are typically created/managed by internal workflows rather than directly by end users.
        - **Allocation intent detail**: `ipType` indicates the type of IP address requested (e.g., IPv4/IPv6 or other model-defined types).

---

# Data Source: intersight_ippool_block_lease
BlockLeases represent an IP address allocation record issued from an IP pool to a specific consuming entity (for example, a server profile). The object acts as the “holder” for an allocated address (or block context) and provides the correlation points needed to track which universe/VRF scope the address belongs to and which entity currently owns it.
#### Purpose
Track and manage IP allocations at the block level so the system can enforce uniqueness, associate the allocation with the correct IP universe and VRF scope, and relate the allocation to the set of underlying IP lease objects.
#### Key Concepts
- **Pool-based allocation record**: Represents an IP address allocated from a pool for use by a specific consumer.
- **Universe-scoped bookkeeping**: The `universe` relationship anchors the allocation in the IP Universe (with cascade on universe deletion) and is used for permission inheritance.
- **VRF-scoped allocation**: `vrf` is `createonly`, ensuring the routing context for the allocation is fixed at creation time.
- **Ownership correlation**: `assignedToEntity` links the allocation to the consuming managed object (e.g., server profile) for traceability and lifecycle management.
- **Lease grouping**: `ipLeases` provides a collection of related `IpLease` objects, enabling a single block-level record to reference one or more concrete lease entries.
- **System-managed lifecycle**: While users can READ, CREATE/UPDATE/DELETE are system API methods, reflecting that allocations are typically created/managed by internal workflows rather than directly by end users.
- **Allocation intent detail**: `ipType` indicates the type of IP address requested (e.g., IPv4/IPv6 or other model-defined types).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_block_lease.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `nr_count`:(int) Count of number of leases requested. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
