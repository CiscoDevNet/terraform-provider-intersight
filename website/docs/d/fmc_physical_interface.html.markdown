---
subcategory: "fmc"
layout: "intersight"
page_title: "Intersight: intersight_fmc_physical_interface"
description: |-
        Physical interface of the FMC device.

---

# Data Source: intersight_fmc_physical_interface
Physical interface of the FMC device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fmc_physical_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Physical interface device Id. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `domain_id`:(string) Physical interface domain Id. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for the physical interface. 
* `physical_interface_id`:(string) The id for the physical interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
