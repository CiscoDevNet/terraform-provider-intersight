---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_local_user_password"
description: |-
        LocalUserPassword type is used for changing local user's password. Caller must send old password in Password field and new password in newPassword field. Intersight will verify the old password and sets the new password if everything is OK. This API must not be used for resetting user's password.

---

# Data Source: intersight_iam_local_user_password
LocalUserPassword type is used for changing local user's password. Caller must send old password in Password field and new password in newPassword field. Intersight will verify the old password and sets the new password if everything is OK. This API must not be used for resetting user's password.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_local_user_password.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `current_password`:(string) User-entered password to be compared to password for change password function. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `initial_password`:(string) Initial password set for the local user for the first time when the local user gets created or when the password gets reset by the Account Administrator. 
* `is_current_password_set`:(bool) Indicates whether the value of the 'currentPassword' property has been set. 
* `is_initial_password_set`:(bool) Indicates whether the value of the 'initialPassword' property has been set. 
* `is_new_password_set`:(bool) Indicates whether the value of the 'newPassword' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `need_password_reset`:(bool) Indicates whether the user should be prompted to reset their password. 
* `new_password`:(string) New password that the user's password should be changed to. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
