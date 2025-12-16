---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_diagnostics"
description: |-
        The health check feature is designed to perform diagnostics on server hardware components.

---

# Data Source: intersight_server_diagnostics
The health check feature is designed to perform diagnostics on server hardware components.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_diagnostics.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action to be performed on the server whether to start or to cancel the diagnostics.* `Start` - Mark the start of the diagnostics on the server.* `Cancel` - Mark the cancellation of the diagnostics on the server.* `Complete` - Mark the completion of the diagnostics on the server. 
* `create_time`:(string) The time when this managed object was created. 
* `diagnostics_type`:(string) Type of diagnostics to be performed on the server hardware components.* `Quick` - Perform fast and limited diagnostics on server hardware components.* `Comprehensive` - Perform slow and extensive diagnostics on server hardware components. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the diagnostics being run. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
