---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_session"
description: |-
  Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
---

# Data Source: intersight_kvm_session
Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_session.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time at which the session ended. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `one_time_password`:(string) Temporary one-time password for vKVM access. 
* `role`:(string) Role of the user who launched the session. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sso_supported`:(bool) Indicates if vKVM SSO is supported on the server. 
* `status`:(string) The status of the session.* `Active` - The session is currently active.* `Ended` - The session has ended normally.* `Terminated` - The session was terminated by an admin. 
* `target_name`:(string) Name of target on which session is initiated. 
* `user_id_or_email`:(string) User ID or E-mail Address of the user who launched the session. 
* `username`:(string) Username used for vKVM access. 
 