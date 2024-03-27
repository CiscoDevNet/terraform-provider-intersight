---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_limits"
description: |-
        The resource limits used to limit resources such as User accounts.

---

# Data Source: intersight_iam_resource_limits
The resource limits used to limit resources such as User accounts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_resource_limits.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allow_api_keys_without_expiry`:(bool) Boolean value used to decide whether API keys that never expire are allowed for the account. This allows creation of API keys which are perpetual which can used for specific applications where rotation of API keys are not feasible. 
* `allow_app_registrations_without_expiry`:(bool) Boolean value used to decide whether App Registration that never expire are allowed for the account. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `max_api_key_expiry`:(int) The maximum expiration period (in seconds) allowed for API keys. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability. 
* `max_app_registration_expiry`:(int) The maximum expiration period (in seconds) allowed for App Registration. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `per_account_user_limit`:(int) The maximum number of users allowed in an account. The default value is 200. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
