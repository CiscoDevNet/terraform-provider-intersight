---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_setting"
description: |-
        Holder for UI Settings such as preference of the user for Session Recording.

---

# Data Source: intersight_iam_user_setting
Holder for UI Settings such as preference of the user for Session Recording.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allow_ui_session_recording`:(bool) UI preference of the user for Session Recording. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id_or_email`:(string) UserID or email as configured in the IdP. 
* `user_unique_identifier`:(string) Unique id of the user used by the identity provider to store the user. 
 
