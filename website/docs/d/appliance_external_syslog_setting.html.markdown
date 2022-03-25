---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_external_syslog_setting"
description: |-
        Configure External Syslog Server.

---

# Data Source: intersight_appliance_external_syslog_setting
Configure External Syslog Server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_external_syslog_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Enable or disable External Syslog Server. 
* `export_nginx`:(bool) Enable or disable exporting of Web Server access logs. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port`:(int) External Syslog Server Port for connection establishment. 
* `protocol`:(string) Protocol used to connect to external syslog server.* `TCP` - External Syslog messages sent over TCP.* `UDP` - External Syslog messages sent over UDP.* `TLS` - Secure External Syslog messages sent over TLS. 
* `server`:(string) External Syslog Server Address, can be IP address or hostname. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
