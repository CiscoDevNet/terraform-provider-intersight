---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_network"
description: |-
        Depicts network configuration used to create a virtual network.

---

# Data Source: intersight_virtualization_virtual_network
Depicts network configuration used to create a virtual network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_virtual_network.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Human readable description about this network. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `infra_network`:(bool) A flag to distinguish if a network belongs to an infrastructure network or a user defined network that guest workloads can use. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmissible unit of data in bytes that can be sent across the network. 
* `name`:(string) Name of the virtual network. Name must be unique. 
* `network_type`:(string) Type of network layer, either L2 or L3.* `unknown` - This network is of an unknown network type.* `L2` - A Layer 2 switching network type. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vlan`:(int) A VLAN id set on the network attachment point. 
* `vswitch`:(string) Name of the virtual switch. 
 
