
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the BuildTaskMeta instance used to created a dynamic workflow. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `src`:(string) Microservice owner for the task in this workflow. 
* `task_type`:(string) The type of the task within this workflow. 
* `workflow_type`:(string) The type for the dynamic workflow. 
