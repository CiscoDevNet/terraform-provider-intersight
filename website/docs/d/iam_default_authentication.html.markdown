---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_default_authentication"
description: |-
        Default authentication method as chosen by the user to login into Intersight Appliance.

---

# Data Source: intersight_iam_default_authentication
Default authentication method as chosen by the user to login into Intersight Appliance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_default_authentication.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_authentication_method`:(string) The default authentication method used to login into Intersight.* `Local` - Local authentication uses credentials stored within the Intersight platform for user access.* `SSO` - SSO authentication uses an external identity provider for user access.* `LDAP` - LDAP authentication uses external LDAP servers for user access. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
