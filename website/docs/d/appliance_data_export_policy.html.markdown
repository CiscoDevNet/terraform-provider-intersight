---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_data_export_policy"
description: |-
        DataExportPolicies define category-based data collection controls for exporting data from an Intersight Appliance to Intersight. The policy is organized as a hierarchy of categories and sub-categories (for example, Global → Inventory → Network/Storage, and TechSupport), allowing administrators to enable or disable export at different levels while preserving a clear parent/child structure.
        #### Purpose
        Provide centralized, hierarchical control over which categories of appliance data are collected and exported to Intersight, supporting governance, privacy/compliance needs, and operational control of data collection behavior.
        #### Key Concepts
        - **Hierarchical category model**: Policies form a tree of configurations where a parent category (e.g., Inventory) can contain sub-configurations (e.g., Network, Storage).
        - **Cascading enable/disable semantics**: Enabling or disabling a category implicitly enables/disables all of its sub-categories (e.g., toggling Inventory also affects Network and Storage).
        - **Primary control switch**: `enable` is the core read/write flag indicating whether data export (collection) is enabled for that category node.
        - **System-managed naming**: `name` is read-only for API consumers but system-writable (`sysapiaccess: readwrite`), indicating naming may be controlled/seeded by the appliance/system.
        - **Account-scoped governance**: Inherits permissions from `account`, and relates to `iam.Account` to scope policy ownership and access.
        - **Parent/child lifecycle handling**: `subConfigs` are children of a given category; deletion cascades to sub-configurations, preserving hierarchy consistency.
        - **Operational access control**: READ is available to admins/read roles and the “Configure Data Collection Settings” privilege; UPDATE is restricted to administrators with configuration privileges.

---

# Data Source: intersight_appliance_data_export_policy
DataExportPolicies define category-based data collection controls for exporting data from an Intersight Appliance to Intersight. The policy is organized as a hierarchy of categories and sub-categories (for example, Global → Inventory → Network/Storage, and TechSupport), allowing administrators to enable or disable export at different levels while preserving a clear parent/child structure.
#### Purpose
Provide centralized, hierarchical control over which categories of appliance data are collected and exported to Intersight, supporting governance, privacy/compliance needs, and operational control of data collection behavior.
#### Key Concepts
- **Hierarchical category model**: Policies form a tree of configurations where a parent category (e.g., Inventory) can contain sub-configurations (e.g., Network, Storage).
- **Cascading enable/disable semantics**: Enabling or disabling a category implicitly enables/disables all of its sub-categories (e.g., toggling Inventory also affects Network and Storage).
- **Primary control switch**: `enable` is the core read/write flag indicating whether data export (collection) is enabled for that category node.
- **System-managed naming**: `name` is read-only for API consumers but system-writable (`sysapiaccess: readwrite`), indicating naming may be controlled/seeded by the appliance/system.
- **Account-scoped governance**: Inherits permissions from `account`, and relates to `iam.Account` to scope policy ownership and access.
- **Parent/child lifecycle handling**: `subConfigs` are children of a given category; deletion cascades to sub-configurations, preserving hierarchy consistency.
- **Operational access control**: READ is available to admins/read roles and the “Configure Data Collection Settings” privilege; UPDATE is restricted to administrators with configuration privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_data_export_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Data Export Policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
