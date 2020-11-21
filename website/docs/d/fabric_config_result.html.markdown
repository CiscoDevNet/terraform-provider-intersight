
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_config_result"
sidebar_current: "docs-intersight-data-source-fabric-config-result"
description: |-
This provides overall state and detailed information for the deploy and validation profile configuration results.
---

# Data Source: intersight_fabric._config_result
This provides overall state and detailed information for the deploy and validation profile configuration results.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
