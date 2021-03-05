---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_schedule_config_policy"
description: |-
  Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
---

# Data Source: intersight_recovery_schedule_config_policy
Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recovery_schedule_config_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 