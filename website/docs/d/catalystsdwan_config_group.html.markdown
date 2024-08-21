---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_config_group"
description: |-
        Details for the Catalyst SDWAN config groups.

---

# Data Source: intersight_catalystsdwan_config_group
Details for the Catalyst SDWAN config groups.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_config_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_group_id`:(string) UUID for the Catalyst SDWAN config group. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The Catalyst SDWAN config group name. 
* `number_of_devices`:(int) The Catalyst SDWAN config group number of devices. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `solution`:(string) The Catalyst SDWAN config group solution. 
 
