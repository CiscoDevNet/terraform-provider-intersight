
---
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_aggregation"
sidebar_current: "docs-intersight-data-source-cond-alarm-aggregation"
description: |-
Object which represents alarm aggregation for a managed end point.
---

# Data Source: intersight_cond._alarm_aggregation
Object which represents alarm aggregation for a managed end point.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `critical_alarms_count`:(int) Count of all alarms with severity Critical, irrespective of acknowledgement status. 
* `health`:(string) Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
* `info_alarms_count`:(int) Count of all alarms with severity Info, irrespective of acknowledgement status. 
* `mo_type`:(string) Managed object type. For example, FI managed object type will be network.Element. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `warning_alarms_count`:(int) Count of all alarms with severity Warning, irrespective of acknowledgement status. 
