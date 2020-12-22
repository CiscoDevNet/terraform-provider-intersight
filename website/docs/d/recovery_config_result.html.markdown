---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_config_result"
description: |-
  Profile configuration (deploy, validation) results with the overall state and detailed result messages.
---

# Data Source: intersight_recovery_config_result
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
