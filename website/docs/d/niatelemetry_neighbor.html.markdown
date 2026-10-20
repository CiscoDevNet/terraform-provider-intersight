---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_neighbor"
description: |-
        A Layer 3 neighbor managed by a Nexus Dashboard instance.

---

# Data Source: intersight_niatelemetry_neighbor
A Layer 3 neighbor managed by a Nexus Dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_neighbor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `asn`:(int) Local autonomous system number for the neighbor. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_name`:(string) Fabric name for the neighbor. 
* `interface_ip`:(string) Interface IP address for the neighbor. 
* `local_ip`:(string) Local IP address for the neighbor. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operational_status`:(string) Operational status of the neighbor. 
* `ospf_area`:(string) OSPF area for the neighbor. 
* `peer_asn`:(int) Peer autonomous system number for the neighbor. 
* `peer_ip`:(string) Peer IP address for the neighbor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_name`:(string) Switch name for the neighbor. 
* `type`:(string) The type of Layer 3 neighbor object. 
* `vrf_name`:(string) VRF name for the neighbor. 
 
