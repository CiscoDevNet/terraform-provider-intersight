---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_iks_license_count"
description: |-
        IksLicenseCounts represent aggregated usage metrics for IKS licensing (for example, counts of devices in a specific IKS tier).
        #### Purpose
        Provides visibility into IKS tier consumption for monitoring and compliance.
        #### Key Concepts
        - **Tier-based aggregation:** Summarizes counts by tier (e.g., Advantage).
        - **Reporting primitive:** Useful for dashboards and licensing status views.
        - **System-maintained values:** Often updated by backend processes and exposed read-only to users.
        - **Account-scoped:** Tied to AccountLicenseData for correct ownership and permissions.

---

# Data Source: intersight_license_iks_license_count
IksLicenseCounts represent aggregated usage metrics for IKS licensing (for example, counts of devices in a specific IKS tier).
#### Purpose
Provides visibility into IKS tier consumption for monitoring and compliance.
#### Key Concepts
- **Tier-based aggregation:** Summarizes counts by tier (e.g., Advantage).
- **Reporting primitive:** Useful for dashboards and licensing status views.
- **System-maintained values:** Often updated by backend processes and exposed read-only to users.
- **Account-scoped:** Tied to AccountLicenseData for correct ownership and permissions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_iks_license_count.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advantage_count`:(int) The total number of devices claimed in the IKS Advantage tier. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
