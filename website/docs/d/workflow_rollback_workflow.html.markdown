
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_rollback_workflow"
sidebar_current: "docs-intersight-data-source-workflow-rollback-workflow"
description: |-
Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
---

# Data Source: intersight_workflow._rollback_workflow
Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) The action of the rollback workflow such as Create and Start.* `None` - If no action is set, then the default value is set to none for the action field.* `Create` - Create rollback workflow data for the execution of the rollback workflow.* `Start` - Start a new execution of the rollback workflow. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `continue_on_task_failure`:(bool) When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `status`:(string) Status of the rollback workflow instance (Created, Running, Completed, Failed).* `None` - If no status is set, then the default value is set none for the status field.* `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback.* `Running` - Status of the rollback workflow when it is in-progress.* `Completed` - Status of the rollback workflow after execution is successful.* `Failed` - Status of the rollback workflow after execution results in failure. 
