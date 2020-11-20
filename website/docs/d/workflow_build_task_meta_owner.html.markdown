
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta_owner"
sidebar_current: "docs-intersight-data-source-workflow-build-task-meta-owner"
description: |-
Contains the list of dynamic workflow types that a microservice supports.
---

# Data Source: intersight_workflow._build_task_meta_owner
Contains the list of dynamic workflow types that a microservice supports.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `service`:(string) The microservice owner responsible for the tasks. 
