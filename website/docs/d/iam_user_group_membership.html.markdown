---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_group_membership"
description: |-
        Persists a time-limited account-user membership established by a successful regular CUI login. C3 login flows consume this membership without refreshing it or re-evaluating token group claims.

---

# Data Source: intersight_iam_user_group_membership
Persists a time-limited account-user membership established by a successful regular CUI login. C3 login flows consume this membership without refreshing it or re-evaluating token group claims.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_group_membership.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiration_time`:(string) Time after which this retained membership cannot authorize a new C3 session. 
* `last_validated_time`:(string) Time at which a regular CUI login most recently validated this UserGroup membership. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
