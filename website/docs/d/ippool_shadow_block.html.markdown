---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_shadow_block"
description: |-
        The ShadowBlocks object represents a block of contiguous IP addresses that are part of a shadow pool.
        #### Purpose
        It manages the specific ranges of IP addresses available for allocation within a shadow pool context, which is created per VRF to track IP usage.
        #### Key Concepts
        - **Range Management:** Defines contiguous blocks of IPv4 or IPv6 addresses.
        - **VRF Association:** Links IP blocks to specific Virtual Routing and Forwarding (VRF) domains.
        - **Pool Integration:** Functions as a core component of shadow pools for structured IP tracking.

---

# Data Source: intersight_ippool_shadow_block
The ShadowBlocks object represents a block of contiguous IP addresses that are part of a shadow pool.
#### Purpose
It manages the specific ranges of IP addresses available for allocation within a shadow pool context, which is created per VRF to track IP usage.
#### Key Concepts
- **Range Management:** Defines contiguous blocks of IPv4 or IPv6 addresses.
- **VRF Association:** Links IP blocks to specific Virtual Routing and Forwarding (VRF) domains.
- **Pool Integration:** Functions as a core component of shadow pools for structured IP tracking.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_shadow_block.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `ip_type`:(string) Type of this IP addresses blocks.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
