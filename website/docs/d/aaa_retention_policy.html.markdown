---
subcategory: "aaa"
layout: "intersight"
page_title: "Intersight: intersight_aaa_retention_policy"
description: |-
        The RetentionPolicy object governs audit log retention policies at the account level, specifying the duration for which audit logs are preserved before automatic deletion.
        #### Purpose
        A RetentionPolicy defines the time frame for audit log retention, ensuring logs are kept for an appropriate duration to meet organizational and regulatory requirements. This facilitates systematic management of log retention, balancing storage needs with compliance obligations.
        #### Key Concepts
        - **Policy Management:** Allows administrators to define and update retention periods, providing flexibility for changing compliance needs.
        - **Automatic Cleanup:** Enables automatic log deletion based on retention periods, reducing manual overhead and ensuring efficient storage use.
        - **Access Control:** Privilege sets govern the creation, reading, and updating of RetentionPolicies, ensuring only authorized users can modify configurations.
        - **Account Association:** Each RetentionPolicy is linked to specific accounts, enabling tailored retention strategies for different organizational units.

---

# Data Source: intersight_aaa_retention_policy
The RetentionPolicy object governs audit log retention policies at the account level, specifying the duration for which audit logs are preserved before automatic deletion.
#### Purpose
A RetentionPolicy defines the time frame for audit log retention, ensuring logs are kept for an appropriate duration to meet organizational and regulatory requirements. This facilitates systematic management of log retention, balancing storage needs with compliance obligations.
#### Key Concepts
- **Policy Management:** Allows administrators to define and update retention periods, providing flexibility for changing compliance needs.
- **Automatic Cleanup:** Enables automatic log deletion based on retention periods, reducing manual overhead and ensuring efficient storage use.
- **Access Control:** Privilege sets govern the creation, reading, and updating of RetentionPolicies, ensuring only authorized users can modify configurations.
- **Account Association:** Each RetentionPolicy is linked to specific accounts, enabling tailored retention strategies for different organizational units.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_aaa_retention_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `retention_period`:(int) The time period in months for audit log retention. Audit logs beyond this period will be automatically deleted. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
