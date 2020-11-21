
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_pending_dynamic_workflow_info"
sidebar_current: "docs-intersight-data-source-workflow-pending-dynamic-workflow-info"
description: |-
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
---

# Data Source: intersight_workflow._pending_dynamic_workflow_info
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name for the pending dynamic workflow. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `src`:(string) The src is workflow owner service. 
* `status`:(string) The current status of the PendingDynamicWorkflowInfo.* `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution.* `Waiting` - Dynamic workflow is in waiting state and not yet started execution. 
* `wait_on_duplicate`:(bool) When set to true workflow engine will wait for a duplicate to finish before starting a new one. 
* `workflow_key`:(string) This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup. 
