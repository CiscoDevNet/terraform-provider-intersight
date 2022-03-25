---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_tfc_organization"
description: |-
        Organizations are a shared space for teams to collaborate on workspaces in Terraform Cloud.

---

# Data Source: intersight_cloud_tfc_organization
Organizations are a shared space for teams to collaborate on workspaces in Terraform Cloud.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_tfc_organization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `agent_ceiling`:(int) The max number of active agents allowed in this organization. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email`:(string) The admin email address associated with a user under this organization. 
* `identity`:(string) The ID of the organization. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the organization. 
* `num_teams`:(int) The number of teams under this organization. 
* `num_users`:(int) The number of users in this organization. 
* `run_ceiling`:(int) The max number of simultaneous runs allowed in this organization. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vcs_providers`:(int) Total number of VCS providers in the organization. 
 
