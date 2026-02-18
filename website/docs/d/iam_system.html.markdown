---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_system"
description: |-
        The System object is the global root within the platform, serving as a container for objects that are shared across all accounts. It centralizes platform-wide security constructs, such as privileges, privilege sets, and roles, as well as the default identity provider and service provider configuration.
        #### Purpose
        System provides a consistent, read-only global context from which account-scoped security and access models are derived and referenced.
        #### Key Concepts
        - **Global Scope:** Hosts objects that are common and reusable across all accounts.
        - **Security Model Anchor:** Aggregates privilege sets, privileges, and roles available platform-wide.
        - **Identity Foundation:** Holds the default identity provider and service provider definitions.
        - **Read-only Access:** Ensures integrity of shared security constructs.

---

# Data Source: intersight_iam_system
The System object is the global root within the platform, serving as a container for objects that are shared across all accounts. It centralizes platform-wide security constructs, such as privileges, privilege sets, and roles, as well as the default identity provider and service provider configuration.
#### Purpose
System provides a consistent, read-only global context from which account-scoped security and access models are derived and referenced.
#### Key Concepts
- **Global Scope:** Hosts objects that are common and reusable across all accounts.
- **Security Model Anchor:** Aggregates privilege sets, privileges, and roles available platform-wide.
- **Identity Foundation:** Holds the default identity provider and service provider definitions.
- **Read-only Access:** Ensures integrity of shared security constructs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_system.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
