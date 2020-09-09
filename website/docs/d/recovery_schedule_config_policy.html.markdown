
---
layout: "intersight"
page_title: "Intersight: intersight_recovery_schedule_config_policy"
sidebar_current: "docs-intersight-data-source-recovery-schedule-config-policy"
description: |-
Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
---

# Data Source: intersight_recovery._schedule_config_policy
Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
