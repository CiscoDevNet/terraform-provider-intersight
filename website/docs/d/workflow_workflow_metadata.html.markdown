
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_metadata"
sidebar_current: "docs-intersight-data-source-workflow-workflow-metadata"
description: |-
Workflow metadata is a collection of properties that are common across all the versions of a workflow.
---

# Data Source: intersight_workflow._workflow_metadata
Workflow metadata is a collection of properties that are common across all the versions of a workflow.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) The description for this workflow. 
* `label`:(string) A user friendly short name to identify the workflow metadata. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this workflow metadata. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
