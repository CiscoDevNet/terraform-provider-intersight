---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_datacenter"
description: |-
  A datacenter object in HyperFlex Application Platform. It is a pre-defined object created internally by the system which acts as a container (logically) for all other objects (Host, VirtualMachine, Volume etc).
---

# Data Source: intersight_hyperflex_hxap_datacenter
A datacenter object in HyperFlex Application Platform. It is a pre-defined object created internally by the system which acts as a container (logically) for all other objects (Host, VirtualMachine, Volume etc).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_datacenter.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 