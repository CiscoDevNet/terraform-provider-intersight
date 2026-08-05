---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_policy"
description: |-
        The IAM EndPointUserPolicy object is a reusable policy that enables the creation and management of local user accounts on endpoints, such as server management controllers.
        #### Purpose
        The purpose of this policy is to standardize local user account configurations across multiple servers. It allows administrators to define password complexity rules, password expiration settings, and associate specific roles and permissions with local users, ensuring consistent and secure access control.
        #### Key Concepts
        - **Centralized User Management:** Provides a single policy to define local user accounts, their roles, and password settings that can be applied to many servers.
        - **Password Properties:** The policy includes a comprehensive set of password controls, such as enforcing strong passwords, setting expiration durations, and maintaining a password history.
        - **Role-Based Access Control:** It works in conjunction with EndPointUser and EndPointUserRole objects to map users to specific roles (e.g., admin, read-only) on the endpoint.
        - **Profile-Based Application:** The policy is attached to a Server Profile to apply the user configurations to the assigned physical server.

---

# Data Source: intersight_iam_end_point_user_policy
The IAM EndPointUserPolicy object is a reusable policy that enables the creation and management of local user accounts on endpoints, such as server management controllers.
#### Purpose
The purpose of this policy is to standardize local user account configurations across multiple servers. It allows administrators to define password complexity rules, password expiration settings, and associate specific roles and permissions with local users, ensuring consistent and secure access control.
#### Key Concepts
- **Centralized User Management:** Provides a single policy to define local user accounts, their roles, and password settings that can be applied to many servers.
- **Password Properties:** The policy includes a comprehensive set of password controls, such as enforcing strong passwords, setting expiration durations, and maintaining a password history.
- **Role-Based Access Control:** It works in conjunction with EndPointUser and EndPointUserRole objects to map users to specific roles (e.g., admin, read-only) on the endpoint.
- **Profile-Based Application:** The policy is attached to a Server Profile to apply the user configurations to the assigned physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_lockout_duration`:(int) Timeout duration specifies the duration (in seconds) after which a locked account is automatically unlocked. - Set to 0 when accountUnlockMode is Manual. - Set a value between 1 and 604800 when accountUnlockMode is Automatic. 
* `account_lockout_threshold`:(int) Set Account Lockout Threshold for endpoint users. 
* `account_moid`:(string) The Account ID for this managed object. 
* `account_unlock_mode`:(string) Account unlock method specifies how the account is unlocked after it is locked: - Manual: Account must be manually unlocked by an administrator. - Automatic: Account unlocks automatically after a timeout duration.* `Automatic` - Set Automatic on the selected end point.* `Manual` - Set Manual on the selected end point. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
