---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta_owner"
description: |-
  Contains the list of dynamic workflow types that a microservice supports.
---

# Data Source: intersight_workflow_build_task_meta_owner
Contains the list of dynamic workflow types that a microservice supports.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_build_task_meta_owner.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `service`:(string) The microservice owner responsible for the tasks. 
 