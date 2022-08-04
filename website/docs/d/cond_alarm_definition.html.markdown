---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm_definition"
description: |-
        The definition of an issue which encompases the criteria for identifying when the issue exists, documentation of the detected issue and the alarms to be raised/cleared by Intersight.

---

# Data Source: intersight_cond_alarm_definition
The definition of an issue which encompases the criteria for identifying when the issue exists, documentation of the detected issue and the alarms to be raised/cleared by Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cond_alarm_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) A description of the issue which is common to all instances of the issue. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An informational display name. 
* `probable_cause`:(string) An explanation of the likely causes of the detected issue. 
* `remediation`:(string) An explanation of the steps to perform to remediate the detected issue. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
