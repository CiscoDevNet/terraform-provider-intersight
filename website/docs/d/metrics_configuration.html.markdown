---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_configuration"
description: |-
        Configurations control account-wide metric collection behavior. It defines whether metric collection is enabled for the account (system-controlled), how newly added resources should be treated, and what platform limits/granularity apply.
        #### Purpose
        Provide a single, account-scoped control point for how metrics are collected across all eligible resources in the account, including default behavior for newly claimed devices.
        #### Key Concepts
        - **Account-wide governance:** Applies settings consistently across all metric-capable resources in the account.
        - **New resource policy:** Controls whether new resources are auto-enabled for metrics collection or require explicit enablement.
        - **Platform-enforced limits:** Exposes system-defined limits such as maximum enabled resources and minimum supported collection granularity.
        - **Advanced ingestion tuning:** Supports an optional user override for max ingestion rate (intended for use with TAC guidance).
        - **Account binding and permissions:** Identity is the account; access is controlled via account-level privileges.

---

# Data Source: intersight_metrics_configuration
Configurations control account-wide metric collection behavior. It defines whether metric collection is enabled for the account (system-controlled), how newly added resources should be treated, and what platform limits/granularity apply.
#### Purpose
Provide a single, account-scoped control point for how metrics are collected across all eligible resources in the account, including default behavior for newly claimed devices.
#### Key Concepts
- **Account-wide governance:** Applies settings consistently across all metric-capable resources in the account.
- **New resource policy:** Controls whether new resources are auto-enabled for metrics collection or require explicit enablement.
- **Platform-enforced limits:** Exposes system-defined limits such as maximum enabled resources and minimum supported collection granularity.
- **Advanced ingestion tuning:** Supports an optional user override for max ingestion rate (intended for use with TAC guidance).
- **Account binding and permissions:** Identity is the account; access is controlled via account-level privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `collect_new_devices`:(string) The behavior of the system when new resources are added, controls whether metric collection are automatically enabled for the new resources.* `AutoEnable` - Automatically enable metric collection for new resources, up to the limit of resource collection.* `Disabled` - Metrics will not be enabled on new resources, to enable collection requires an explicit user enable. 
* `collection_granularity`:(string) The current supported collection granularity by the system, defined as the lowest granularity supported, with the actual granularity per resource determined by the license tier of the resource. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Enables metric collection for the account, if disabled metrics will be stopped for all resources in the account. 
* `limit`:(int) The total number of resources that can be enabled for metric collection in this account. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_max_ingestion_bytes`:(int) User override of the max ingestion rate for metrics. Provided as an advanced option to override the default limits. Option should only be used in coordination with TAC engineers. 
 
