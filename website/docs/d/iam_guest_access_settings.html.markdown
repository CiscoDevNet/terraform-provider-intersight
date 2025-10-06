---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_guest_access_settings"
description: |-
        Stores account level settings for guest user access.

---

# Data Source: intersight_iam_guest_access_settings
Stores account level settings for guest user access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_guest_access_settings.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `max_guest_access_link_shelf_life`:(int) Stores the maximum duration (in seconds) during which guest access link remains valid and accessible. It is the maximum value that is set  at the account level which account admin can configure. Any guest access link that is set with expiration time beyond this property will be disallowed. The default value is set to 604800 seconds (7 days). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `override_ip_access_restriction`:(bool) Stores an option for Account Admin to override IP Access Restriction if it is enabled in the Account. This option is used to disable IP Access restrictions for guest users logging in to the account, while restriction is enforced for other normal users (who are authenticated via SAML or LDAP). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
