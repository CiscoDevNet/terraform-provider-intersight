---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_profile"
description: |-
        A profile specifying configuration settings for a chassis.

---

# Data Source: intersight_chassis_profile
A profile specifying configuration settings for a chassis.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `chassis_assignment_mode`:(string) Source of the chassis assigned to the Chassis Profile. Values can be Static or None. Static is used if a chassis is attached directly to a Chassis Profile. None is used if no chassis is attached to a Chassis Profile. Slot or Serial pre-assignment is also considered to be None as it is different form of Assign Later.* `Static` - Chassis is directly assigned to chassis profile using assign chassis.* `None` - No chassis is assigned to the chassis profile. 
* `chassis_pre_assign_by_serial`:(string) Serial number of the chassis that would be assigned to this pre-assigned Chassis Profile. It can be any string that adheres to the following constraints:It should start and end with an alphanumeric character.It cannot be more than 20 characters. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_platform`:(string) The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `user_label`:(string) User label assigned to the chassis profile. 
 
