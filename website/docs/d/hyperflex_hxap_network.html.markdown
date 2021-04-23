---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_network"
description: |-
  A HyperFlex Application Platform network attachment entity that a VM can attached to.
---

# Data Source: intersight_hyperflex_hxap_network
A HyperFlex Application Platform network attachment entity that a VM can attached to.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_network.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The internally generated identity of network. This entity cannot manipulated by users. It aids in uniquely identifying the network object. For VMware, this is MOR (managed object reference). 
* `infra_network`:(bool) A flag to distinguish if a network belongs to a HxAp infrastructure network or a user defined network that guest workloads can use. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The MTU size of the network. 
* `name`:(string) User-provided name to identify the portgroup. 
* `network_type`:(string) Network attachment type, only \ L2\  is available now.* `unknown` - This network is of an unknown network type.* `L2` - A Layer 2 switching network type. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vlan`:(int) A vlan id set on the network attachment point. 
 