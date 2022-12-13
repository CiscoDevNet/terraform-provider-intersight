---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_reservation"
description: |-
        The WWN reservation object, used to hold reserved identities.

---

# Data Source: intersight_fcpool_reservation
The WWN reservation object, used to hold reserved identities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_reservation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the allocation for the identity in the reservation either static or dynamic (i.e. via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `id_purpose`:(string) Purpose of this WWN ID. Purpose can be WWPN or WWNN. 
* `identity`:(string) WWN ID that needs to be reserved. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
