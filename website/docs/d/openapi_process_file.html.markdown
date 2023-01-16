---
subcategory: "openapi"
layout: "intersight"
page_title: "Intersight: intersight_openapi_process_file"
description: |-
        Validates the OpenAPI specification file. On successful validation, it persists information about the available APIs. This information can be used to create tasks in the cloud orchestrator.

---

# Data Source: intersight_openapi_process_file
Validates the OpenAPI specification file. On successful validation, it persists information about the available APIs. This information can be used to create tasks in the cloud orchestrator.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_openapi_process_file.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `failure_reason`:(string) An error message for when download, validation or processing of OpenAPI Specification fails. 
* `file_download_status`:(string) Status of the internal OpenAPI specification fetch operation* `none` - Indicates the default status* `InProgress` - Indicates that operation is in progress* `Completed` - Indicates that the operation is complete* `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details. 
* `file_processing_status`:(string) Status of the OpenAPI specification processing operation. The OpenAPI specification file is processed to create APIMethod objects.* `none` - Indicates the default status* `InProgress` - Indicates that operation is in progress* `Completed` - Indicates that the operation is complete* `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details. 
* `file_validation_status`:(string) Status of the OpenAPI specification validation operation.* `none` - Indicates the default status* `InProgress` - Indicates that operation is in progress* `Completed` - Indicates that the operation is complete* `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `spec_file_path`:(string) The location of the previously uploaded OpenAPI specification file. 
 
