
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_catalog"
sidebar_current: "docs-intersight-data-source-workflow-catalog"
description: |-
A catalog of workflow related objects such as workflow and task definitions. Each user account will have a local workflow catalog where account users can store their private workflow and task definitions.
Cisco provides validated workflows and tasks to Intersight users via shared catalogs. Intersight users will be able to read, run these workflows and tasks within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
---

# Data Source: intersight_workflow._catalog
A catalog of workflow related objects such as workflow and task definitions. Each user account will have a local workflow catalog where account users can store their private workflow and task definitions.
Cisco provides validated workflows and tasks to Intersight users via shared catalogs. Intersight users will be able to read, run these workflows and tasks within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
