---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_lease"
description: |-
        Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.

---

# Data Source: intersight_fcpool_lease
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_lease.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `has_duplicate`:(bool) HasDuplicate represents if there are other pools in which this id exists. 
* `migrate`:(bool) The migration capability is applicable only for dynamic lease requests and it works in conjunction with  preferred ID. If there is an existing dynamic or static lease that matches the preferred ID, that existing  lease will be migrated to the current pool. That means the existing lease will be deleted and a new lease  will be created in the pool. If there is a reservation exists that matches with preferred ID, that  reservation will be kept as is and next available ID from the pool will be leased. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pool_purpose`:(string) Purpose of this WWN pool. 
* `preferred_wwn_id`:(string) The preferred WWN ID address can be specified only for dynamic lease requests. Intersight will make its best  effort to allocate that WWN ID address if it is available in the pool. If the specified preferred WWN ID address  is not in the range of the pool or if it is already leased or reserved, then the next available WWN ID address  from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease  request will fail if it specifies the preferred WWN ID address property. When the preferred WWN ID address  property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be  replaced by the new lease. Migration also supported only for dynamic lease requests. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `wwn_id`:(string) WWN ID allocated for pool based allocation. 
 
