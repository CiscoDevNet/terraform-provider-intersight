---
subcategory: "sdaaci"
layout: "intersight"
page_title: "Intersight: intersight_sdaaci_connection_detail"
description: |-
        Peer connection details for each connection.

---

# Data Source: intersight_sdaaci_connection_detail
Peer connection details for each connection.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdaaci_connection_detail.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of this connection between two peers. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_pool`:(string) Id of the ip pool configured for this connection. 
* `layer3_handoff_id`:(string) Id of layer 3 handoff configured between a border node and a border leaf. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `peer_ainterface`:(string) Interface id configured on Peer A. 
* `peer_aip_address`:(string) The IP Address of the device used as the local peer. 
* `peer_atype`:(string) Type of device used as Peer A for this peer connection. 
* `peer_binterface`:(string) Interface id configured on Peer B. 
* `peer_bip_address`:(string) The IP Address of the device used as the remote peer. 
* `peer_btype`:(string) Type of device used as Peer B for this peer connection. 
* `peera`:(string) First peer of the connection. 
* `peerb`:(string) Second Peer of the connection. 
* `router_id`:(string) Router id defined for this peer connection. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Connection status between the peers.* `NotConnected` - Connection Status NotConnected.* `Connected` - Connection Status Connected. 
 
