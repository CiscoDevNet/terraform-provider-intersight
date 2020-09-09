
---
layout: "intersight"
page_title: "Intersight: intersight_recovery_config_result"
sidebar_current: "docs-intersight-data-source-recovery-config-result"
description: |-
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
---

# Data Source: intersight_recovery._config_result
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
