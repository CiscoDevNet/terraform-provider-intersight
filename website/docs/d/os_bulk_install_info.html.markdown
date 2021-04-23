---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_bulk_install_info"
description: |-
  This MO models the CSV file content which the user uploaded for OS installation. As part of the handler, necessary filed
in the model can be populated along with respective validation.
---

# Data Source: intersight_os_bulk_install_info
This MO models the CSV file content which the user uploaded for OS installation. As part of the handler, necessary filed
in the model can be populated along with respective validation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_bulk_install_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_content`:(string) The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections.The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holdsparameters which are specific to each server level data. 
* `is_file_content_set`:(bool) Indicates whether the value of the 'fileContent' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the CSV file, which holds the OS install parameters. 
* `oper_state`:(string) Denotes if the operating is pending, in_progress, completed_ok, completed_error.* `Pending` - The initial value of the OperStatus.* `InProgress` - The OperStatus value will be InProgress during execution.* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 