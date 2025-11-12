---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_qualifier"
description: |-
        User Qualifier holds list of guest user details configured for guest user group type, which then can be used for guest user login.

---

# Data Source: intersight_iam_user_qualifier
User Qualifier holds list of guest user details configured for guest user group type, which then can be used for guest user login.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_qualifier.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_notification_time`:(string) Holds the information on when last email notification was sent to the guest users. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `notification_scope`:(string) Defines the scope of email notifications for guest users. It determines which guest users  should receive email notifications about account access details. Options include notifying all users or only  newly added users.* `All` - Email Notification is sent to all users.* `New` - Email Notification is sent to newly added users. 
* `notify_guest_users`:(bool) NotifyGuestUsers holds information on whether guest users have been notified about the guest access information. If set to true, all guest users will receive a email notification about the details of guest access link and instructions. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
