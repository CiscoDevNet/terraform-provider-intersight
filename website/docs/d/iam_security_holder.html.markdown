---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_security_holder"
description: |-
        SecurityHolder aggregates account-wide and organization-scoped permissions, serving as the nexus for effective authorization state across the accounts resource hierarchy.
        #### Purpose
        SecurityHolder centralizes derived permissions and related controls like IP access management for consistent policy enforcement.
        #### Key Concepts
        - **Aggregated Permissions:** Collects effective permissions per resource (organization).
        - **Policy Container:** Anchors resource-permission mappings for authorization checks.
        - **Access Governance:** Hosts IP-based access management settings for the account.
        - **Account Integration:** Bound to the accounts overall security structure.

---

# Data Source: intersight_iam_security_holder
SecurityHolder aggregates account-wide and organization-scoped permissions, serving as the nexus for effective authorization state across the accounts resource hierarchy.
#### Purpose
SecurityHolder centralizes derived permissions and related controls like IP access management for consistent policy enforcement.
#### Key Concepts
- **Aggregated Permissions:** Collects effective permissions per resource (organization).
- **Policy Container:** Anchors resource-permission mappings for authorization checks.
- **Access Governance:** Hosts IP-based access management settings for the account.
- **Account Integration:** Bound to the accounts overall security structure.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_security_holder.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
