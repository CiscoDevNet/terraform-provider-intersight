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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `identity`:(string) Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter. 
