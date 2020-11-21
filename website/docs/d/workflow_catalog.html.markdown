
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
