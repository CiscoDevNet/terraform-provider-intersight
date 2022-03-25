---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_session"
description: |-
        The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.

---

# Data Source: intersight_iam_session
The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_session.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time at which the session ended. 
* `expiration`:(string) Expiration time for the session. 
* `failed_logins`:(int) Failed logins since last login for admin user. 
* `idle_time_expiration`:(string) Idle time expiration for the session. 
* `last_login_client`:(string) The client address from which last login is initiated. 
* `last_login_time`:(string) The last login time for user. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `role`:(string) Role of the user who launched the session. 
* `session_id`:(string) Session token shared with the user agent which is used to identify the user session when API requests are received to perform authorization. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the session.* `Active` - The session is currently active.* `Ended` - The session has ended normally.* `Terminated` - The session was terminated by an admin. 
* `user_id_or_email`:(string) User ID or E-mail Address of the user who launched the session. 
 
