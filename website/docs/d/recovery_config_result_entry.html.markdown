---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_config_result_entry"
description: |-
  An entry that describes the result of a Backup Profile state on the end device.
---

# Data Source: intersight_recovery_config_result_entry
An entry that describes the result of a Backup Profile state on the end device.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `completed_time`:(string) The completed time of the task in installer. 
* `message`:(string) Localized message based on the locale setting of the user's context. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owner_id`:(string) The identifier of the object that owns the result message.The owner ID is used to correlate a given result entry to a task or entity. For example, a config resultentry that describes the result of a workflow task may have the task's instance ID as the owner. 
* `state`:(string) Values  -- Ok, Ok-with-warning, Errored. 
* `type`:(string) Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config. 
