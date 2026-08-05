---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_batch_api_executor"
description: |-
        The BatchApiExecutor object is a key component within the workflow system that facilitates the execution of multiple API requests as part of a single task execution. It is designed to handle batch operations efficiently, allowing users to define and execute a series of API calls in a cohesive manner.
        #### Purpose
        The BatchApiExecutor serves as a mechanism to associate and manage a list of API requests that need to be executed together. This provides the ability to define API tasks using the API request body and a response parser specification, thus streamlining the execution process within a workflow.
        #### Key Concepts
        - **Batch Execution:** Enables the grouping and execution of multiple API requests as a single task, enhancing operational efficiency.
        - **Access Control:** Utilizes privilege sets to manage who can read, update, create, and delete batch API tasks, ensuring security and proper governance.
        - **Error Handling:** Integrates with error response handlers to provide robust mechanisms for parsing and managing error responses from API calls.

---

# Data Source: intersight_workflow_batch_api_executor
The BatchApiExecutor object is a key component within the workflow system that facilitates the execution of multiple API requests as part of a single task execution. It is designed to handle batch operations efficiently, allowing users to define and execute a series of API calls in a cohesive manner.
#### Purpose
The BatchApiExecutor serves as a mechanism to associate and manage a list of API requests that need to be executed together. This provides the ability to define API tasks using the API request body and a response parser specification, thus streamlining the execution process within a workflow.
#### Key Concepts
- **Batch Execution:** Enables the grouping and execution of multiple API requests as a single task, enhancing operational efficiency.
- **Access Control:** Utilizes privilege sets to manage who can read, update, create, and delete batch API tasks, ensuring security and proper governance.
- **Error Handling:** Integrates with error response handlers to provide robust mechanisms for parsing and managing error responses from API calls.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_batch_api_executor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed description of the batch APIs. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the batch API task. 
* `retry_from_failed_api`:(bool) Flag indicating if the retry task should from the failed API or the first API in the batch execution; default value is false. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_on_condition`:(string) Optional skip expression allowing the batch API executor to skip task execution when the provided Go template expression evaluates to true. If not specified, the API will always be executed. 
 
