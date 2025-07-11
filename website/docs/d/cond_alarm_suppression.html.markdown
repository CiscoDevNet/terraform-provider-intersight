---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_suppression"
description: |-
        Intersight allows alarms to be suppressed at different Intersight entity MO instances. This model captures the suppression information such as which alarms to be suppressed at which entity MO.
        An Intersight entity MO can have suppressions created directly at the entity MO level or at any of the entity MO ancestors. The resultant alarms to be suppressed at an entity MO level will be based on the combination of suppressions created for that entity and its ancestors. Currently only server maintenance is supported.

---

# Data Source: intersight_cond_alarm_suppression
Intersight allows alarms to be suppressed at different Intersight entity MO instances. This model captures the suppression information such as which alarms to be suppressed at which entity MO.
An Intersight entity MO can have suppressions created directly at the entity MO level or at any of the entity MO ancestors. The resultant alarms to be suppressed at an entity MO level will be based on the combination of suppressions created for that entity and its ancestors. Currently only server maintenance is supported.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_alarm_suppression.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User given description on why the suppression is enabled at this entity. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name that identifies the alarm suppression. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
