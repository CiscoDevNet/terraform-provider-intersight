
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
