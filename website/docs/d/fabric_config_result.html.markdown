---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_config_result"
description: |-
  This provides overall state and detailed information for the deploy and validation profile configuration results.
---

# Data Source: intersight_fabric_config_result
This provides overall state and detailed information for the deploy and validation profile configuration results.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_config_result.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_stage`:(string) The current running stage of the configuration or workflow. 
* `config_state`:(string) Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `validation_state`:(string) Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. 
 