---
subcategory: "workspace"
layout: "intersight"
page_title: "Intersight: intersight_workspace_folder"
description: |-
        Folder contains list of predefined assets like Workflows, Service items, Widgets, etc. it enables customers to make uses of these readily available predefined assets.

---

# Data Source: intersight_workspace_folder
Folder contains list of predefined assets like Workflows, Service items, Widgets, etc. it enables customers to make uses of these readily available predefined assets.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workspace_folder.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `archived`:(bool) It is to define if folder is archived or not. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The UserID or email who created this folder. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The UserID or email who last modified this folder. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this folder. You can have multiple versions of the folder with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
