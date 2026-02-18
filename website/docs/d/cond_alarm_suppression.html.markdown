---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_suppression"
description: |-
        The AlarmSuppression is for selective suppression of alarms across various managed objects (MOs). It allows users to define which alarms should be hidden or suppressed, ensuring that only relevant alerts are visible and actionable.
        #### Purpose
        AlarmSuppression provides a mechanism to control and manage alarms by suppressing them at specific entity levels. This capability is essential for maintaining focus on critical alerts and minimizing unnecessary noise in alarm systems.
        #### Key Concepts
        - **Entity-Based Suppression:** Alarms can be suppressed directly at the MO level.
        - **Privilege-Based Access:** Access to create, update, and delete suppressions is governed by privilege sets, ensuring that only authorized users can modify alarm suppression settings.
        - **Operational Flexibility:** Suppressions can be tailored to specific server maintenance needs, allowing users to focus on critical alerts during maintenance windows.

---

# Data Source: intersight_cond_alarm_suppression
The AlarmSuppression is for selective suppression of alarms across various managed objects (MOs). It allows users to define which alarms should be hidden or suppressed, ensuring that only relevant alerts are visible and actionable.
#### Purpose
AlarmSuppression provides a mechanism to control and manage alarms by suppressing them at specific entity levels. This capability is essential for maintaining focus on critical alerts and minimizing unnecessary noise in alarm systems.
#### Key Concepts
- **Entity-Based Suppression:** Alarms can be suppressed directly at the MO level.
- **Privilege-Based Access:** Access to create, update, and delete suppressions is governed by privilege sets, ensuring that only authorized users can modify alarm suppression settings.
- **Operational Flexibility:** Suppressions can be tailored to specific server maintenance needs, allowing users to focus on critical alerts during maintenance windows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_alarm_suppression.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User given description on why the suppression is enabled at this entity. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Indicates whether the suppression is enabled by the user or not. The user should be able to toggle this between true and false.The property is set to true when the suppression is created. The user can set this to false to disable the suppression.The suppression rule should be active only if both systemEnabled and enabled are true. 
* `end_date`:(string) The end date for this alarm suppression rule. The date must follow the RFC 3339 format for date and time representation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name that identifies the alarm suppression. 
* `odata_filter_internal`:(string) Odata filter string managed internally. It is built by combining all the rules. 
* `rules_operator`:(string) Operation that binds all the different rules together.* `All` - All is an AND condition applied against the individual conditions.* `Any` - Any is an OR condition applied against the individual conditions. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_date`:(string) The start date for enabling this alarm suppression rule. The date must followthe RFC 3339 format for date and time representation. If this date more than60 seconds in the past, the suppression rule will be rejected. If the date iswithin 60 seconds of the present time (plus or minus), the suppression will bestarted immediately. Otherwise, the suppression will be scheduled to start atthe requested time. 
 
