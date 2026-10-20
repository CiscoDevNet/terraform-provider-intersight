---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_role_inventory"
description: |-
        The IAM EndPointUserRole object is a mapping object that links a local user to their assigned roles and password within the context of an EndPointUserPolicy.
        #### Purpose
        The IAM EndPointUserRole objects purpose is to bind together three key pieces of information: the user, their password, and their permissions. It serves as the mechanism for granting specific roles (like admin or read-only) to an EndPointUser and securely storing the encrypted password for that user's account on the endpoint.
        #### Key Concepts
        - **User-Role-Password Mapping:** It connects an EndPointUser object with a collection of EndPointRole objects and contains the password for that user.
        - **Secure Password Handling:** The password property is write-only and encrypted, ensuring credentials are not exposed after being set.
        - **Component of a Policy:** It exists as part of the endPointUserRoles collection within an EndPointUserPolicy and is essential for defining a complete user account.
        - **Account Status:** Includes an enabled flag to control whether the user account is active on the endpoint.

---

# Data Source: intersight_iam_end_point_user_role_inventory
The IAM EndPointUserRole object is a mapping object that links a local user to their assigned roles and password within the context of an EndPointUserPolicy.
#### Purpose
The IAM EndPointUserRole objects purpose is to bind together three key pieces of information: the user, their password, and their permissions. It serves as the mechanism for granting specific roles (like admin or read-only) to an EndPointUser and securely storing the encrypted password for that user's account on the endpoint.
#### Key Concepts
- **User-Role-Password Mapping:** It connects an EndPointUser object with a collection of EndPointRole objects and contains the password for that user.
- **Secure Password Handling:** The password property is write-only and encrypted, ensuring credentials are not exposed after being set.
- **Component of a Policy:** It exists as part of the endPointUserRoles collection within an EndPointUserPolicy and is essential for defining a complete user account.
- **Account Status:** Includes an enabled flag to control whether the user account is active on the endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user_role_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `account_type_user_defined`:(bool) Allows to choose custom account types for the endpoint user. 
* `change_password`:(bool) Denotes whether password has changed. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Enables the user account on the endpoint. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) The password must have a minimum of 8 and a maximum of 127 characters. For servers with IPMI user role enabled, the maximum length is limited to 20 characters. When strong password is enabled, must satisfy below requirements: A. The password must not contain the User's Name. B. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &, *, -, _, +, =). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
