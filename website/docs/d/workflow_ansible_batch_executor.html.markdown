---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_ansible_batch_executor"
description: |-
        Intersight allows generic tasks to be created by taking the executor request
        body and a response parser specification in the form of content.Grammar object.
        Ansible Batch associates the list of Ansible commands executed as SSH requests as part of single
        task execution. Each SSH request takes the Ansible command to execute and a response parser
        specification based off text to extract fields of interest.

---

# Data Source: intersight_workflow_ansible_batch_executor
Intersight allows generic tasks to be created by taking the executor request
body and a response parser specification in the form of content.Grammar object.
Ansible Batch associates the list of Ansible commands executed as SSH requests as part of single
task execution. Each SSH request takes the Ansible command to execute and a response parser
specification based off text to extract fields of interest.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_ansible_batch_executor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) A detailed description about the batch APIs. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the batch API task. 
* `retry_from_failed_api`:(bool) When an execution of a nth API in the Batch fails,Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry.By default the value is set to false. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_on_condition`:(string) The skip expression, if provided, allows the batch API executor to skip thetask execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed. 
 
