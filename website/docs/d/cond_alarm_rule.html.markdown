---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_rule"
description: |-
        A container object that aggregates a set of ThresholdDefinition objects. This can be assigned to one or more devices, triggering the creation of the AlarmDefinition which enables centralized management and consistent application of ThresholdDefinition objects across multiple resources.

---

# Data Source: intersight_cond_alarm_rule
A container object that aggregates a set of ThresholdDefinition objects. This can be assigned to one or more devices, triggering the creation of the AlarmDefinition which enables centralized management and consistent application of ThresholdDefinition objects across multiple resources.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_alarm_rule.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Informative description of AlarmRule. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name assigned by the user to AlarmRule. This user-defined name acts as the identity field, ensuring that AlarmRule is distinctly identifiable within the account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Controls the behavior of alarm processing depending upon the state. If Enabled, which is also the default behavior, the alarm is evaluated for the device based on the conditions specified in the ThresholdDefinition objects attached to it. If Disabled or SystemDisabled, alarm is not evaluated for the device and the existing alarms raised against the device is cleared.* `Enabled` - User initiated action which is also the default action that enables alarm evaluation for any condition that meets the criteria.* `Disabled` - User initiated action that disables alarm evaluation temporarily and clears the existing alarms.* `SystemDisabled` - System initiated action that disables alarm evaluation temporarily and clears the existing alarms once alarm limit per alarm rule is reached. 
* `system_disabled_time`:(string) The time at which AlarmRule object is system-disabled. 
 
