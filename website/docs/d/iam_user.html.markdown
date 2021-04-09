---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user"
description: |-
  The Intersight account user.
---

# Data Source: intersight_iam_user
The Intersight account user.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip_address`:(string) IP address from which the user last logged in to Intersight. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email`:(string) Email of the user. Users are added to Intersight using the email configured in the IdP. 
* `first_name`:(string) First name of the user. This field is populated from the IdP attributes received after authentication. 
* `last_login_time`:(string) Last successful login time for user. 
* `last_name`:(string) Last name of the user. This field is populated from the IdP attributes received after authentication. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name as configured in the IdP. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id_or_email`:(string) UserID or email as configured in the IdP. 
* `user_type`:(string) Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic. 
* `user_unique_identifier`:(string) Unique id of the user used by the identity provider to store the user. 
 