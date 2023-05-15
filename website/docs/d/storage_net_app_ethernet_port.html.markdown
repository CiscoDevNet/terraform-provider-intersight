---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_ethernet_port"
description: |-
        Ethernet port is a port on a node in a storage array.

---

# Data Source: intersight_storage_net_app_ethernet_port
Ethernet port is a port on a node in a storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_ethernet_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `broadcast_domain_name`:(string) Name of the broadcast domain, scoped to its IPspace. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(string) Status of port to determine if its enabled or not. 
* `mac_address`:(string) MAC address of the port available in storage array. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit of the physical port available in storage array. 
* `name`:(string) Name of the port available in storage array. 
* `node_name`:(string) The node name for the port. 
* `port_state`:(string) State of the port available in storage array.* `Down` - An inactive port is listed as Down.* `Up` - An active port is listed as Up.* `Degraded` - An active port that is Up but unhealthy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(int) Operational speed of port measured. 
* `state`:(string) State of the port available in storage array.* `down` - An inactive port is listed as Down.* `up` - An active port is listed as Up.* `present` - An active port is listed as present. 
* `type`:(string) Type of the port available in storage array.* `LAG` - Storage port of type lag.* `physical` - LIFs can be configured directly on physical ports.* `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port. 
* `uuid`:(string) Universally unique identifier of the physical port. 
 
