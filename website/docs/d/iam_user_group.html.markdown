---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_group"
description: |-
        User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.

---

# Data Source: intersight_iam_user_group
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_activation_time`:(string) AccessActivationTime indicates the activation time for the guest user's access to the Account.  Before this time, if guest user tries to login to the account, access the account will be denied. 
* `access_expiry_time`:(string) AccessExpiryTime indicates the expiration time for the guest user's access to the Account. Its value can only be  assigned a date that falls within the range determined by the maximum expiration time configured for the  API entries. The AccessExpiry date can be edited to be earlier or later. 
* `access_link`:(string) AccessLink using which the guest user uses to log in to Intersight. 
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `group_type`:(string) Group type determines the type of groups that is being associated with users. By default, Default User group will be used for associating dynamic user login. If the value of the User Group is set to guest, then this type of user group will be used for guest user login.* `Default` - Default User Group Type used for dynamic users login.* `Guest` - Guest User Group type used for guest users login. 
* `instruction`:(string) Instruction property holds detailed guidance and information intended for individuals  accessing the system as guest users. It holds the information to assist guests in navigating the platform,  understanding policies, and performing necessary actions to ensure a seamless and secure user experience. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the user group which the dynamic/or guest user belongs to. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `unique_reference_id`:(string) A random mixed character string which is unique per user groups. UniqueReferenceId is used as key for identifying the guest user groups. 
 
