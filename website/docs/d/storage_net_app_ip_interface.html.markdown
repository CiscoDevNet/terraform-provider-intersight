---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_ip_interface"
description: |-
        NetApp IP interface is a logical interface.

---

# Data Source: intersight_storage_net_app_ip_interface
NetApp IP interface is a logical interface.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_ip_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(string) IP interface is enabled or not. 
* `home_node`:(string) Name of home node of IP interface. 
* `home_port`:(string) Name of home port of IP interface. 
* `ip_address`:(string) The IP address of interface. 
* `ip_family`:(string) IP address family of interface.* `IPv4` - IP address family type is IPv4.* `IPv6` - IP address family type is IP6. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the IP interface. 
* `netmask`:(string) The netmask of the interface. 
* `service_policy_name`:(string) Service policy name of IP interface. 
* `service_policy_uuid`:(string) Service policy UUID of IP interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the IP interface.* `down` - An inactive port is listed as Down.* `up` - An active port is listed as Up.* `present` - An active port is listed as present. 
* `uuid`:(string) Uuid of NetApp IP Interface. 
 
