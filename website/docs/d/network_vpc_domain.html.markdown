---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_vpc_domain"
description: |-
        Concrete class for VPC domain configured on a network device. VPC (Virtual Port Channel) domain is used to connect two different switches logically to a single switch.

---

# Data Source: intersight_network_vpc_domain
Concrete class for VPC domain configured on a network device. VPC (Virtual Port Channel) domain is used to connect two different switches logically to a single switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_vpc_domain.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_recovery_status`:(string) Auto recovery status of the virtual port channel domain. 
* `consistency_status`:(string) Consistency status of the virtual port channel domain. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dual_active_excluded_vlans`:(int) Dual Active Excluded VLANs of the virtual port channel domain. 
* `keep_alive_status`:(string) Keep alive status of the virtual port channel domain. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `peer_status`:(string) Peer status of the virtual port channel domain. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) Role of the virtual port channel domain. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vpc_domain_id`:(int) Identity of the virtual port channel domain. 
* `vpcs_configured_count`:(int) Number of VPCs configured on the virtual port channel domain. 
 
