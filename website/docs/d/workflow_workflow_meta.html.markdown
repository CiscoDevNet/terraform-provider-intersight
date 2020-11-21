
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_meta"
sidebar_current: "docs-intersight-data-source-workflow-workflow-meta"
description: |-
Contains a workflow definition which is a sequence of tasks to execute.
---

# Data Source: intersight_workflow._workflow_meta
Contains a workflow definition which is a sequence of tasks to execute.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) The description for the workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name given to the workflow. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `retryable`:(bool) When true, this workflow can be retried for 2 weeks since the last modification of the workflow. 
* `schema_version`:(int) The Conductor schema version that decides what attribute can be supported. 
* `src`:(string) The src is workflow owner service. 
* `type`:(string) The type of workflow definition.* `SystemDefined` - System defined workflow definition.* `UserDefined` - User defined workflow definition.* `Dynamic` - Dynamically defined workflow definition. 
* `version`:(int) The version for the workflow so we can support multiple versions for the same workflow name. 
* `wait_on_duplicate`:(bool) Parameter decides if workflows will wait for a duplicate to finish before starting a new one. 
