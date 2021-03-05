---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_pending_dynamic_workflow_info"
description: |-
  Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
---

# Data Source: intersight_workflow_pending_dynamic_workflow_info
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_pending_dynamic_workflow_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name for the pending dynamic workflow. 
* `src`:(string) The src is workflow owner service. 
* `status`:(string) The current status of the PendingDynamicWorkflowInfo.* `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution.* `Waiting` - Dynamic workflow is in waiting state and not yet started execution. 
* `wait_on_duplicate`:(bool) When set to true workflow engine will wait for a duplicate to finish before starting a new one. 
* `workflow_key`:(string) This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup. 
 