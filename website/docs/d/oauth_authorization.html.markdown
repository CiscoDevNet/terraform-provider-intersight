---
subcategory: "oauth"
layout: "intersight"
page_title: "Intersight: intersight_oauth_authorization"
description: |-
  User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user. It is used by Intersight Appliance to support resource owner grant type.
---

# Data Source: intersight_oauth_authorization
User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user. It is used by Intersight Appliance to support resource owner grant type.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_oauth_authorization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `api_type`:(string) Type of OAuth Api. For example, Smart-licensing-API.* `Unknown` - Unknown is the default API type.* `SmartLicensing-API` - Smart licensing API type. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool) Indicates whether the value of the 'userId' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id`:(string) The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
 