---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_account"
description: |-
        The Account object defines the organizational framework within which users and resources are managed, serving as the central hub for access and permissions.
        #### Purpose
        An Account encapsulates the structure and permissions associated with a specific organizational unit, providing a secure environment for managing users and resources.
        #### Key Concepts
        - **Access Management:** Governs user access and permissions, ensuring appropriate roles are assigned to users within the account.
        - **Resource Organization:** Structures resources under the account, facilitating management and oversight of assets.
        - **Security Configuration:** Enforces security policies and configurations, including IP access management and session limits.
        - **Identity Providers:** Integrates with identity providers to streamline authentication and user management, enhancing security and user experience.

---

# Data Source: intersight_iam_account
The Account object defines the organizational framework within which users and resources are managed, serving as the central hub for access and permissions.
#### Purpose
An Account encapsulates the structure and permissions associated with a specific organizational unit, providing a secure environment for managing users and resources.
#### Key Concepts
- **Access Management:** Governs user access and permissions, ensuring appropriate roles are assigned to users within the account.
- **Resource Organization:** Structures resources under the account, facilitating management and oversight of assets.
- **Security Configuration:** Enforces security policies and configurations, including IP access management and session limits.
- **Identity Providers:** Integrates with identity providers to streamline authentication and user management, enhancing security and user experience.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_account.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_identifier`:(string) External identifier for the account, used for integration with external identity systems. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Intersight account. By default, name is same as the MoID of the account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `single_admin_lockout`:(bool) Indicates if the account is prone to lockout as it has only a single Account Administrator. An account is prone to lockout if it has only one configured Account Administrator and no user groups configured that can grant Account Administrator role to dynamic users. 
* `status`:(string) Status of the account. To activate the Intersight account, claim a device to the account. 
 
