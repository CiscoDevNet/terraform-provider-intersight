---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_backup_profile"
description: |-
        The BackupProfile object serves as the cornerstone for managing backup configurations on endpoints. It defines the parameters and conditions under which backups are conducted, supporting both scheduled and on-demand operations.
        #### Purpose
        The BackupProfile object defines and manages backup configurations for endpoints. It allows setting up backup parameters, including scheduling and execution details, and enables or disables backup operations as needed to ensure systematic data protection, integrity, and availability.
        #### Key Concepts
        - **Flexibility:** Allows for both on-demand and scheduled backups, catering to diverse operational requirements.
        - **Versioning and Validation:** Supports robust validation mechanisms to ensure the accuracy and reliability of backup configurations.
        - **Relationship Management:** Links to device registrations and configuration results to maintain a comprehensive backup strategy.

---

# Data Source: intersight_recovery_backup_profile
The BackupProfile object serves as the cornerstone for managing backup configurations on endpoints. It defines the parameters and conditions under which backups are conducted, supporting both scheduled and on-demand operations.
#### Purpose
The BackupProfile object defines and manages backup configurations for endpoints. It allows setting up backup parameters, including scheduling and execution details, and enables or disables backup operations as needed to ensure systematic data protection, integrity, and availability.
#### Key Concepts
- **Flexibility:** Allows for both on-demand and scheduled backups, catering to diverse operational requirements.
- **Versioning and Validation:** Supports robust validation mechanisms to ensure the accuracy and reliability of backup configurations.
- **Relationship Management:** Links to device registrations and configuration results to maintain a comprehensive backup strategy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recovery_backup_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Enables/Disables the schedule on the endpoint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
 
