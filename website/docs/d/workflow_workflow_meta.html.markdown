
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) The description for the workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name given to the workflow. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `retryable`:(bool) When true, this workflow can be retried for 2 weeks since the last modification of the workflow. 
* `schema_version`:(int) The Conductor schema version that decides what attribute can be supported. 
* `src`:(string) The src is workflow owner service. 
* `type`:(string) The type of workflow definition.* `SystemDefined` - System defined workflow definition.* `UserDefined` - User defined workflow definition.* `Dynamic` - Dynamically defined workflow definition. 
* `version`:(int) The version for the workflow so we can support multiple versions for the same workflow name. 
* `wait_on_duplicate`:(bool) Parameter decides if workflows will wait for a duplicate to finish before starting a new one. 
