---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_aggregation"
description: |-
        The AlarmAggregation object is designed to represent aggregated alarm data for a managed endpoint within Intersight. It offers a consolidated view of alarm counts and severity levels, supporting efficient monitoring and management of endpoint health.
        #### Purpose
        AlarmAggregation provides a summary of alarm data, aiding users in understanding the overall health and alarm status of managed endpoints. It is essential for assessing the impact of alarms on system performance and decision-making.
        #### Key Concepts
        - **Summary of Severity:** Aggregates alarm counts based on severity types (Critical, Warning, Info), giving users a clear view of alarm distribution.
        - **Endpoint Health Assessment:** Evaluates and reports the health status of endpoints based on aggregated alarm data, supporting proactive management.
        - **Relationship-Driven:** Tied to alarm aggregation sources within the managed object hierarchy, ensuring accurate and relevant aggregation data.
        - **Read-Only Access:** Designed for safe consumption, allowing users to integrate and assess alarm data without modifying it directly.

---

# Data Source: intersight_cond_alarm_aggregation
The AlarmAggregation object is designed to represent aggregated alarm data for a managed endpoint within Intersight. It offers a consolidated view of alarm counts and severity levels, supporting efficient monitoring and management of endpoint health.
#### Purpose
AlarmAggregation provides a summary of alarm data, aiding users in understanding the overall health and alarm status of managed endpoints. It is essential for assessing the impact of alarms on system performance and decision-making.
#### Key Concepts
- **Summary of Severity:** Aggregates alarm counts based on severity types (Critical, Warning, Info), giving users a clear view of alarm distribution.
- **Endpoint Health Assessment:** Evaluates and reports the health status of endpoints based on aggregated alarm data, supporting proactive management.
- **Relationship-Driven:** Tied to alarm aggregation sources within the managed object hierarchy, ensuring accurate and relevant aggregation data.
- **Read-Only Access:** Designed for safe consumption, allowing users to integrate and assess alarm data without modifying it directly.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_alarm_aggregation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `critical_alarms_count`:(int) Count of all alarms with severity Critical, irrespective of acknowledgement status. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health`:(string) Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
* `info_alarms_count`:(int) Count of all alarms with severity Info, irrespective of acknowledgement status. 
* `mo_type`:(string) Managed object type. For example, FI managed object type will be network.Element. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `warning_alarms_count`:(int) Count of all alarms with severity Warning, irrespective of acknowledgement status. 
 
