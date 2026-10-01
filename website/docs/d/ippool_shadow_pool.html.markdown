---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_shadow_pool"
description: |-
        The ShadowPools object acts as a tracking object created on behalf of an IP pool, scoped specifically to a VRF.
        #### Purpose
        It provides a mechanism to monitor IP pool usage, availability, and address distribution within individual routing contexts.
        #### Key Concepts
        - **VRF-Specific Tracking:** Maintains pool state and usage metrics per VRF.
        - **Resource Monitoring:** Tracks the size and assigned count of IPv4/IPv6 addresses within the shadow context.

---

# Data Source: intersight_ippool_shadow_pool
The ShadowPools object acts as a tracking object created on behalf of an IP pool, scoped specifically to a VRF.
#### Purpose
It provides a mechanism to monitor IP pool usage, availability, and address distribution within individual routing contexts.
#### Key Concepts
- **VRF-Specific Tracking:** Maintains pool state and usage metrics per VRF.
- **Resource Monitoring:** Tracks the size and assigned count of IPv4/IPv6 addresses within the shadow context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_shadow_pool.<custom_name>.results[i].<propertyname>`.
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
* `v4_assigned`:(int) Number of IPv4 addresses currently in use. 
* `v4_size`:(int) Number of IPv4 addresses in this pool. 
* `v6_assigned`:(int) Number of IPv6 addresses currently in use. 
* `v6_size`:(int) Number of IPv6 addresses in this pool. 
 
