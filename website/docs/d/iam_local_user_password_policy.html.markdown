---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_local_user_password_policy"
description: |-
  Configuration for LocalUserPasswordPolicy.
---

# Data Source: intersight_iam_local_user_password_policy
Configuration for LocalUserPasswordPolicy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_local_user_password_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `min_char_difference`:(int) Minimum number of characters different from previous password. 
* `min_days_between_password_change`:(int) Minimum Days allowed between password change. 
* `min_length_password`:(int) Minimum length of password. 
* `min_lower_case`:(int) Minimum number of required lower case characters. 
* `min_numeric`:(int) Minimum number of required numeric characters. 
* `min_special_char`:(int) Minimum number of required special characters. 
* `min_upper_case`:(int) Minimum number of required upper case characters. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_previous_passwords_disallowed`:(int) Number of previous passwords disallowed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 