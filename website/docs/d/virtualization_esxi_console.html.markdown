---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_esxi_console"
description: |-
        Vmware Esxi Console Session that provides request validation and Session Management.

---

# Data Source: intersight_virtualization_esxi_console
Vmware Esxi Console Session that provides request validation and Session Management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_esxi_console.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time at which the session ended. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `role`:(string) Role of the user who launched the session. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the session.* `Active` - The session is currently active.* `Ended` - The session has ended normally.* `Terminated` - The session was terminated by an admin. 
* `stream_id`:(string) The stream ID of the host console session opened. 
* `target_name`:(string) Name of target on which session is initiated. 
* `user_id_or_email`:(string) User ID or E-mail Address of the user who launched the session. 
 
