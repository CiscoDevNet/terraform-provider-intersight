---
subcategory: "oauth"
layout: "intersight"
page_title: "Intersight: intersight_oauth_access_token"
description: |-
        API access token associated with a given account.

---

# Data Source: intersight_oauth_access_token
API access token associated with a given account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_oauth_access_token.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `api_type`:(string) OAuth API type (e.g., Smart licensing API).* `Unknown` - Unknown serves as the default API type.* `SmartLicensing-API` - Smart licensing API type.* `CommerceEstimate-API` - Commerce Estimate API type. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiry`:(string) Access token expiration date and time. 
* `issuer`:(string) OAuth access token issuer. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `refresh_expiry`:(string) Refresh token expiration date and time. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `token_owner`:(string) MOID of the access token owner. 
 
