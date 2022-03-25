---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_tfc_workspace"
description: |-
        Terraform workspace which represents running infrastructure managed by Terraform.

---

# Data Source: intersight_cloud_tfc_workspace
Terraform workspace which represents running infrastructure managed by Terraform.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_tfc_workspace.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `agent_pool_id`:(string) The agent pool associated with this workspace. 
* `apply_method`:(bool) Whether or not Terraform Cloud should automatically apply a successful Terraform plan. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execution_mode`:(string) Indicates where the Terraform cloud should execute the runs. 
* `identity`:(string) The unique id for this workspace. 
* `last_run_status`:(string) The status of the last executed run in this workspace. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the workspace. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
