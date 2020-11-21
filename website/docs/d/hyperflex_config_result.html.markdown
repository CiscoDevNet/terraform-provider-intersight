
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_config_result"
sidebar_current: "docs-intersight-data-source-hyperflex-config-result"
description: |-
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
---

# Data Source: intersight_hyperflex._config_result
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `config_progress`:(string) The progress percentage of the running configuration or workflow. 
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `duration`:(string) The duration of the running configuration or workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `start_time`:(string) The start time of the configuration or workflow. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
