---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_ssh_batch_executor"
description: |-
        The SshBatchExecutor object is a core component of Intersight’s workflow system, specialized in handling the batch execution of SSH commands. It is designed to execute a sequence of SSH requests as part of a single task, improving automation and efficiency in remote command execution.
        #### Purpose
        The SshBatchExecutor simplifies the execution of SSH commands by grouping them into batches, enabling streamlined and secure operations across remote systems.
        #### Key Concepts
        - **Batch Execution:** Supports executing multiple SSH commands in a single batch to optimize workflow efficiency.
        - **SSH Session:** Establishes an SSH session from an Intersight-connected endpoint to a remote server, allowing multiple SSH operations to be executed sequentially over the same session.
        - **File Transfer:** Supports transferring files from an Intersight-connected device to a remote server.

---

# Data Source: intersight_workflow_ssh_batch_executor
The SshBatchExecutor object is a core component of Intersight’s workflow system, specialized in handling the batch execution of SSH commands. It is designed to execute a sequence of SSH requests as part of a single task, improving automation and efficiency in remote command execution.
#### Purpose
The SshBatchExecutor simplifies the execution of SSH commands by grouping them into batches, enabling streamlined and secure operations across remote systems.
#### Key Concepts
- **Batch Execution:** Supports executing multiple SSH commands in a single batch to optimize workflow efficiency.
- **SSH Session:** Establishes an SSH session from an Intersight-connected endpoint to a remote server, allowing multiple SSH operations to be executed sequentially over the same session.
- **File Transfer:** Supports transferring files from an Intersight-connected device to a remote server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_ssh_batch_executor.<custom_name>.results[i].<propertyname>`.
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
 
