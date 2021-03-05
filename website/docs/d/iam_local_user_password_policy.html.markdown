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
* `min_char_difference`:(int) Minimum number of characters different from previous password. 
* `min_days_between_password_change`:(int) Minimum Days allowed between password change. 
* `min_length_password`:(int) Minimum length of password. 
* `min_lower_case`:(int) Minimum number of required lower case characters. 
* `min_numeric`:(int) Minimum number of required numeric characters. 
* `min_special_char`:(int) Minimum number of required special characters. 
* `min_upper_case`:(int) Minimum number of required upper case characters. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_previous_passwords_disallowed`:(int) Number of previous passwords disallowed. 
 