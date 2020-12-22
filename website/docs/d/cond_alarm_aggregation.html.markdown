---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_aggregation"
description: |-
  Object which represents alarm aggregation for a managed end point.
---

# Data Source: intersight_cond_alarm_aggregation
Object which represents alarm aggregation for a managed end point.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `critical_alarms_count`:(int) Count of all alarms with severity Critical, irrespective of acknowledgement status. 
* `health`:(string) Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
* `info_alarms_count`:(int) Count of all alarms with severity Info, irrespective of acknowledgement status. 
* `mo_type`:(string) Managed object type. For example, FI managed object type will be network.Element. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `warning_alarms_count`:(int) Count of all alarms with severity Warning, irrespective of acknowledgement status. 
