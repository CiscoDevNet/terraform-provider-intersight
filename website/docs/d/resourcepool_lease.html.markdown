---
subcategory: "resourcepool"
layout: "intersight"
page_title: "Intersight: intersight_resourcepool_lease"
description: |-
        Lease API invoked by passing resource pool, lease API will reserve or un-reserve the resource from the pool.

---

# Data Source: intersight_resourcepool_lease
Lease API invoked by passing resource pool, lease API will reserve or un-reserve the resource from the pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resourcepool_lease.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `feature`:(string) Lease opertion applied for the feature. 
* `has_duplicate`:(bool) HasDuplicate represents if there are other pools in which this id exists. 
* `is_exclusive_at_assigned_entity`:(bool) Indicates whether a lease allocation is exclusive based on the Assigned Entity, if the AssignedEntity holds any lease then not allowed to create new lease later. 
* `migrate`:(bool) The migration capability is applicable only for dynamic lease requests and it works in conjunction with  preferred ID. If there is an existing dynamic or static lease that matches the preferred ID, that existing  lease will be migrated to the current pool. That means the existing lease will be deleted and a new lease  will be created in the pool. If there is a reservation exists that matches with preferred ID, that  reservation will be kept as is and next available ID from the pool will be leased. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `resource_type`:(string) The type of resource present in the pool, such as 'server' can be a RackUnit or Blade.* `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.* `Mixed` - Resource Pool holds the resources like Rack Server, Blade or Chassis.* `None` - The resource cannot consider for Resource Pool. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
