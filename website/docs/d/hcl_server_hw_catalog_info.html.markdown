---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_server_hw_catalog_info"
description: |-
        Server hardware catalog information for a particular server model, used for server BOM validation.

---

# Data Source: intersight_hcl_server_hw_catalog_info
Server hardware catalog information for a particular server model, used for server BOM validation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_server_hw_catalog_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `display_model`:(string) Display model of the server hardware. In many cases, the model string used in the catalog will be the hardware Cisco PID and the Intersight model string is a more user-friendly string with vendor information in it. This will be the user-friendly modal string to be used in Intersight. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the server hardware from the catalog file. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of specific tag, required to choose the correct datatype while reading the value. 
 
