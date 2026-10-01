---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_vpc_peer"
description: |-
        VpcPeers represent the vPC peer-link configuration on a network device. The peer link is the critical inter-switch port-channel that carries control traffic and, depending on design, some data-plane traffic between the two vPC peers.
        #### Purpose
        Provide visibility into peer-link identity and operational state so operators can validate peer-link health and diagnose vPC peer connectivity issues.
        #### Key Concepts
        - **Domain association**: `vpcDomainId` ties the peer link to its vPC domain.
        - **Peer-link identity**: `vpcPeerId` identifies the peer-link record; `portChannel`/`portChannelId` identify the peer-link port-channel.
        - **Operational state**: `operationalState` indicates whether the peer link is functioning.
        - **Relationship to interface model**: `etherPortChannel` links to the underlying port-channel for interface-level inspection.
        - **Device association**: `registeredDevice` links the peer-link configuration to the specific device.

---

# Data Source: intersight_network_vpc_peer
VpcPeers represent the vPC peer-link configuration on a network device. The peer link is the critical inter-switch port-channel that carries control traffic and, depending on design, some data-plane traffic between the two vPC peers.
#### Purpose
Provide visibility into peer-link identity and operational state so operators can validate peer-link health and diagnose vPC peer connectivity issues.
#### Key Concepts
- **Domain association**: `vpcDomainId` ties the peer link to its vPC domain.
- **Peer-link identity**: `vpcPeerId` identifies the peer-link record; `portChannel`/`portChannelId` identify the peer-link port-channel.
- **Operational state**: `operationalState` indicates whether the peer link is functioning.
- **Relationship to interface model**: `etherPortChannel` links to the underlying port-channel for interface-level inspection.
- **Device association**: `registeredDevice` links the peer-link configuration to the specific device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_vpc_peer.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operational_state`:(string) Operational state of the virtual port channel. 
* `port_channel`:(string) Name of the virtual port channel. 
* `port_channel_id`:(int) Port channel identity of the virtual port channel. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vpc_domain_id`:(int) Identity of the virtual port channel. 
* `vpc_peer_id`:(int) Identity of the virtual port channel. 
 
