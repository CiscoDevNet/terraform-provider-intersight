---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_ext_eth_interface"
description: |-
  Physical port of a virtual interface card.
---

# Data Source: intersight_adapter_ext_eth_interface
Physical port of a virtual interface card.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_ext_eth_interface.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin configured state of an External Ethernet Interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ep_dn`:(string) Endpoint Config DN of an External Ethernet Interface. 
* `ext_eth_interface_id`:(string) Unique Identifier for an External Ethernet Interface within the adapter object. 
* `interface_type`:(string) Type of an External Ethernet Interface. 
* `mac_address`:(string) MAC address of an External Ethernet Interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of an Interface. 
* `peer_aggr_port_id`:(int) Peer Aggregate Port Id attached to an External Ethernet Interface. 
* `peer_dn`:(string) DN of peer end-point attached to an External Ethernet Interface. 
* `peer_port_id`:(int) Peer Port Id attached to an External Ethernet Interface. 
* `peer_slot_id`:(int) Peer Slot Id attached to an External Ethernet Interface. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) SwitchId attached to an External Ethernet Interface. 
 