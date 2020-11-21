
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta"
sidebar_current: "docs-intersight-data-source-workflow-build-task-meta"
description: |-
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
---

# Data Source: intersight_workflow._build_task_meta
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the BuildTaskMeta instance used to created a dynamic workflow. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `src`:(string) Microservice owner for the task in this workflow. 
* `task_type`:(string) The type of the task within this workflow. 
* `workflow_type`:(string) The type for the dynamic workflow. 
