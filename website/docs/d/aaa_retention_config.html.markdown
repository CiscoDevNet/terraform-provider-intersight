---
subcategory: "aaa"
layout: "intersight"
page_title: "Intersight: intersight_aaa_retention_config"
description: |-
        ### Overview
        The RetentionConfig object provides default configurations for audit log retention, offering a standardized approach to managing retention periods in the absence of specific policies.
        #### Purpose
        RetentionConfig serves as a baseline configuration for audit log retention, ensuring logs are preserved for a defined period by default when no explicit retention policy exists. This helps maintain consistent log management practices across the system.
        #### Key Concepts
        - **Default Settings** - Establishes default retention periods, ensuring systematic log retention even without custom policies.
        - **System Ownership** - Managed by the system to guarantee reliable and consistent application of retention settings across all accounts.
        - **Read-Only Access** - Designed for safe consumption, ensuring default configurations are applied without unauthorized modifications.
        - **Policy Integration** - Complements account-specific RetentionPolicies, harmonizing log retention across the system.

---

# Data Source: intersight_aaa_retention_config
### Overview
The RetentionConfig object provides default configurations for audit log retention, offering a standardized approach to managing retention periods in the absence of specific policies.
#### Purpose
RetentionConfig serves as a baseline configuration for audit log retention, ensuring logs are preserved for a defined period by default when no explicit retention policy exists. This helps maintain consistent log management practices across the system.
#### Key Concepts
- **Default Settings** - Establishes default retention periods, ensuring systematic log retention even without custom policies.
- **System Ownership** - Managed by the system to guarantee reliable and consistent application of retention settings across all accounts.
- **Read-Only Access** - Designed for safe consumption, ensuring default configurations are applied without unauthorized modifications.
- **Policy Integration** - Complements account-specific RetentionPolicies, harmonizing log retention across the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_aaa_retention_config.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `retention_period`:(int) The default retention period in months for audit log retention for accounts without a retention policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
