---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_definition"
description: |-
  Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
---

# Data Source: intersight_workflow_workflow_definition
Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `default_version`:(bool) When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version. 
* `description`:(string) The description for this workflow. 
* `label`:(string) A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `license_entitlement`:(string) License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `max_task_count`:(int) The maximum number of tasks that can be executed on this workflow. 
* `max_worker_task_count`:(int) The maximum number of external (worker) tasks that can be executed on this workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `nr_version`:(int) The version of the workflow to support multiple versions. 
