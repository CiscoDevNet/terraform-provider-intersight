---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_ntp_server"
description: |-
  External NTP time servers ONTAP uses for time adjustment and correction.
---

# Data Source: intersight_storage_net_app_ntp_server
External NTP time servers ONTAP uses for time adjustment and correction.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_ntp_server.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `authentication_enabled`:(bool) Indicates whether or not NTP symmetric authentication is enabled. 
* `authentication_key_id`:(string) NTP symmetric authentication key identifier or index number (ID). 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server`:(string) NTP server host name, IPv4, or IPv6 address. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) NTP protocol version for server. Valid versions are 3, 4, or auto.* `none` - Default unknown NTP protocol version.* `3` - NTP protocol version is 3.* `4` - NTP protocol version is 4.* `auto` - NTP protocol version is auto. 
 