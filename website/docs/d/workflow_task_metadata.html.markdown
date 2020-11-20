
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_metadata"
sidebar_current: "docs-intersight-data-source-workflow-task-metadata"
description: |-
Task details is a collection of properties that are common across all the versions of a task.
---

# Data Source: intersight_workflow._task_metadata
Task details is a collection of properties that are common across all the versions of a task.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) The task metadata description to describe what this task will do when executed. 
* `label`:(string) A user friendly short name to identify the task metadata. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the task metadata. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \ Generic\  if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
