---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_variable"
description: |-
        Variables are user-defined entities that can be shared across workflows. They allow users to set a value once and then reference it from different workflows within the same scope. The variables can be of any type that is supported by the workflow system.

---

# Data Source: intersight_workflow_variable
Variables are user-defined entities that can be shared across workflows. They allow users to set a value once and then reference it from different workflows within the same scope. The variables can be of any type that is supported by the workflow system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_variable.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the environment variable. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the environment variable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
