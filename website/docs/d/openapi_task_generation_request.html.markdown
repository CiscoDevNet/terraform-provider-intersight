---
subcategory: "openapi"
layout: "intersight"
page_title: "Intersight: intersight_openapi_task_generation_request"
description: |-
        Creates a request which has information about the tasks that need to be created from the previously uploaded OpenAPI specification file. This object internally triggers a workflow that creates tasks which can be used in workflows.

---

# Data Source: intersight_openapi_task_generation_request
Creates a request which has information about the tasks that need to be created from the previously uploaded OpenAPI specification file. This object internally triggers a workflow that creates tasks which can be used in workflows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_openapi_task_generation_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `endpoint_type`:(string) Indicates if target endpoint is external or internal. An endpoint is internal if the target is an Intersight resource. For instance, configuring an intersight object using a Task.* `External` - Denotes that the target endpoint is an external API endpoint* `Internal` - Denotes that the target endpoint is a Intersight API endpoint 
* `is_validate_request`:(bool) When this value is set to true, the task name validations happen and provides the task validation status against each of the selected API on the property selectedApis When this value is set to false, an internal workflow to generate the tasks is submitted,  conflicting task names are created with an incremented version. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protocol`:(string) Specifies the REST API protocol being used, which can be one of HTTP or HTTPS.* `HTTPS` - Denotes that the API call uses the HTTPS protocol type* `HTTP` - Denotes that the API call uses the HTTP protocol type 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Depicts the status of the task creation request.* `none` - Indicates the default status* `InProgress` - Request has been picked up for generating tasks from the OpenAPI Specification file* `Completed` - All the tasks from the request have been created* `Failed` - There were failures in generating one or more tasks in the request 
* `task_prefix`:(string) Optional string that can be prefixed to the name of created tasks. 
* `url`:(string) Specifies the URL of the endpoint that the created task communicates with. It is defaulted to intersight.com for internal endpoints. 
 
