---
subcategory: "resourcepool"
layout: "intersight"
page_title: "Intersight: intersight_resourcepool_reservation"
description: |-
        The resource reservation object, used to hold reserved resources for a pool. It captures the reservation identity (for example, a serial), allocation type, and the pool reference so the reservation survives pool membership changes and can be reconciled during import/restore.

---

# Data Source: intersight_resourcepool_reservation
The resource reservation object, used to hold reserved resources for a pool. It captures the reservation identity (for example, a serial), allocation type, and the pool reference so the reservation survives pool membership changes and can be reconciled during import/restore.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resourcepool_reservation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the allocation for the identity in the reservation either static or dynamic (i.e. via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The type and serial number of the resource for display. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `resource_serial`:(string) The serial number of the resource. 
* `resource_type`:(string) The MO type of the resource being reserved. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
