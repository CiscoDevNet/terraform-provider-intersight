---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta"
description: |-
  Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
---

# Data Source: intersight_workflow_build_task_meta
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_build_task_meta.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the BuildTaskMeta instance used to created a dynamic workflow. 
* `src`:(string) Microservice owner for the task in this workflow. 
* `task_type`:(string) The type of the task within this workflow. 
* `workflow_type`:(string) The type for the dynamic workflow. 
 