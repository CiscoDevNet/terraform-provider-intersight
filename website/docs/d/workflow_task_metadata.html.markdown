---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_metadata"
description: |-
  Task details is a collection of properties that are common across all the versions of a task.
---

# Data Source: intersight_workflow_task_metadata
Task details is a collection of properties that are common across all the versions of a task.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_task_metadata.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) The task metadata description to describe what this task will do when executed. 
* `label`:(string) A user friendly short name to identify the task metadata. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the task metadata. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \ Generic\  if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume. 
 