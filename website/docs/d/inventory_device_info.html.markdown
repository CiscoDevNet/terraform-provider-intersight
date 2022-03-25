---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_device_info"
description: |-
        Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.

---

# Data Source: intersight_inventory_device_info
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_device_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_state`:(string) Configuration state of server profile config context. 
* `control_action`:(string) Control action of server profile config context. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_state`:(string) Error state of server profile config context. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of server profile config context. 
* `profile_mo_id`:(string) Server profile MO ID of the server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
