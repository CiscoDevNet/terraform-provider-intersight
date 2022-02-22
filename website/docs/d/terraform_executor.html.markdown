---
subcategory: "terraform"
layout: "intersight"
page_title: "Intersight: intersight_terraform_executor"
description: |-
  Executor is a ManagedObject, aka MO. This API is used to execute terraform scripts.
---

# Data Source: intersight_terraform_executor
Executor is a ManagedObject, aka MO. This API is used to execute terraform scripts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_terraform_executor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `command`:(string) Command to be executed during update operation. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operation`:(string) Enum indicates what operation is being done.* `Create` - Creating a Terraform resource.* `Update` - Updating a Terraform resource.* `Delete` - Deleting a Terraform resource. 
* `platform_type`:(string) The Platform type used in conjunction with 'sourceFolderPath' and 'sourceFolderName' determines unique path for a Terraform workflow. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_folder_name`:(string) Folder Name where Terraform workflows are stored. 
* `source_folder_path`:(string) Relative folder Path where 'sourceFolderName' is located. 
* `source_location`:(string) Flag indicates whether workflow is internal/external. 
* `status`:(string) Status of the terraform execution. 
* `task_id`:(string) TaskId of a pontem workflow is same as the MO. 
 