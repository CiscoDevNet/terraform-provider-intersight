---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_config_result_entry"
description: |-
  The profile configuration (deploy, validation) results details information.
---

# Data Source: intersight_chassis_config_result_entry
The profile configuration (deploy, validation) results details information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_config_result_entry.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `completed_time`:(string) The completed time of the task in installer. 
* `message`:(string) Localized message based on the locale setting of the user's context. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owner_id`:(string) The identifier of the object that owns the result message.The owner ID is used to correlate a given result entry to a task or entity. For example, a config resultentry that describes the result of a workflow task may have the task's instance ID as the owner. 
* `state`:(string) Values  -- Ok, Ok-with-warning, Errored. 
* `type`:(string) Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config. 
 