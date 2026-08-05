---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_operation_setting"
description: |-
        The ConfigOperationSetting object represents account-level settings that support configuration backup and restore operations. It centralizes operational configuration required for consistent and secure handling of backup/restore flows.
        #### Purpose
        ConfigOperationSetting exists to hold account-scoped operational settings used by the system to perform backup and restore reliably, including internal keying material required for secure processing.
        #### Key Concepts
        - **Account-Level Configuration:** Provides a single place to manage operational settings for all backup/restore activity within an account.
        - **Operational Consistency:** Ensures backup/restore executions behave predictably across runs by using consistent account-scoped settings.
        - **Governed Access:** Restricted to administrative roles to preserve integrity of backup and restore security posture.

---

# Data Source: intersight_mgmt_config_operation_setting
The ConfigOperationSetting object represents account-level settings that support configuration backup and restore operations. It centralizes operational configuration required for consistent and secure handling of backup/restore flows.
#### Purpose
ConfigOperationSetting exists to hold account-scoped operational settings used by the system to perform backup and restore reliably, including internal keying material required for secure processing.
#### Key Concepts
- **Account-Level Configuration:** Provides a single place to manage operational settings for all backup/restore activity within an account.
- **Operational Consistency:** Ensures backup/restore executions behave predictably across runs by using consistent account-scoped settings.
- **Governed Access:** Restricted to administrative roles to preserve integrity of backup and restore security posture.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_operation_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_account_key_set`:(bool) Indicates whether the value of the 'accountKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
