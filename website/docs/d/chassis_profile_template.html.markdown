---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_profile_template"
description: |-
        The ProfileTemplate object for chassis profiles defines reusable templates for chassis configurations, supporting rapid deployment and consistent configuration across multiple chassis.
        #### Purpose
        Chassis ProfileTemplate streamlines the creation, management, and propagation of standardized chassis configurations within an organization.
        #### Key Concepts
        - **Template-Based Deployment:** Accelerates chassis onboarding and ensures consistent application of best practices.
        - **Centralized Management:** Templates can be updated to affect all derived chassis profiles, simplifying large-scale configuration changes.
        - **Lifecycle Integration:** Supports derivation and update workflows, integrating with bulk operations for efficient management.

---

# Data Source: intersight_chassis_profile_template
The ProfileTemplate object for chassis profiles defines reusable templates for chassis configurations, supporting rapid deployment and consistent configuration across multiple chassis.
#### Purpose
Chassis ProfileTemplate streamlines the creation, management, and propagation of standardized chassis configurations within an organization.
#### Key Concepts
- **Template-Based Deployment:** Accelerates chassis onboarding and ensures consistent application of best practices.
- **Centralized Management:** Templates can be updated to affect all derived chassis profiles, simplifying large-scale configuration changes.
- **Lifecycle Integration:** Supports derivation and update workflows, integrating with bulk operations for efficient management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_profile_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_platform`:(string) The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `usage`:(int) The count of the chassis profiles derived from the template. 
 
