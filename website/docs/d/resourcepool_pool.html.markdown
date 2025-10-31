---
subcategory: "resourcepool"
layout: "intersight"
page_title: "Intersight: intersight_resourcepool_pool"
description: |-
        Pool represents a collection of resource. The resource can be any MO which has PoolResource meta enabled. The resource in the pool can be reserved or unreserved by using Lease API, reserved/unreserved resources can be used in the entities like server profiles.

---

# Data Source: intersight_resourcepool_pool
Pool represents a collection of resource. The resource can be any MO which has PoolResource meta enabled. The resource in the pool can be reserved or unreserved by using Lease API, reserved/unreserved resources can be used in the entities like server profiles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resourcepool_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The pool is evaluated for resources with associated policies based on action. This action will help users to re-sync the resources for a pool.* `None` - The pool will not be considered for evaluation.* `ReEvaluate` - The resources in the pool will be re-evaluated against the server pool qualification associated with it. 
* `assigned`:(int) Number of IDs that are currently assigned (in use). 
* `assignment_order`:(string) Property is deprecated. Sequential is the only assignment order supported.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `pool_type`:(string) The resource management type in the pool, it can be either static or dynamic.* `Static` - The resources in the pool will not be changed until user manually update it.* `Dynamic` - The resources in the pool will be updated dynamically based on the condition.* `Hybrid` - The resources in the pool can be added by the user statically or dynamically, based on the matching conditions of the qualification policy. If the pool contains both statically added resources and resources added based on the qualification policy, the pool type can be classified as hybrid. 
* `reserved`:(int) Number of IDs that are currently reserved (and not in use). 
* `resource_type`:(string) The type of the resource present in the pool, example 'server' its combination of RackUnit and Blade.* `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.* `Mixed` - Resource Pool holds the resources like Rack Server, Blade or Chassis.* `None` - The resource cannot consider for Resource Pool. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Total number of identifiers in this pool. 
 
