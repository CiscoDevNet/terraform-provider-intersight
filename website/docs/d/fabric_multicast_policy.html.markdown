---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_multicast_policy"
description: |-
  A policy to configure Multicast settings for all the Virtual LAN networks.
---

# Data Source: intersight_fabric_multicast_policy
A policy to configure Multicast settings for all the Virtual LAN networks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_multicast_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `querier_ip_address`:(string) Used to define the IGMP Querier IP address. 
* `querier_ip_address_peer`:(string) Used to define the IGMP Querier IP address of the peer switch. 
* `querier_state`:(string) Administrative state of the IGMP Querier for this VLAN.* `Disabled` - IGMP Querier Disabled State.* `Enabled` - IGMP Querier Enabled State. 
* `snooping_state`:(string) Administrative state of the IGMP Snooping for this VLAN.* `Enabled` - IGMP Snooping Enabled State.* `Disabled` - IGMP Snooping Disabled State. 
 