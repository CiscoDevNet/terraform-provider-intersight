---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_account"
description: |-
        The Intersight Account used to access Intersight.

---

# Data Source: intersight_iam_account
The Intersight Account used to access Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_account.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Intersight account. By default, name is same as the MoID of the account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `single_admin_lockout`:(bool) Indicates if the account is prone to lockout as it has only a single Account Administrator. An account is prone to lockout if it has only one configured Account Administrator and no user groups configured that can grant Account Administrator role to dynamic users. 
* `status`:(string) Status of the account. To activate the Intersight account, claim a device to the account. 
 
