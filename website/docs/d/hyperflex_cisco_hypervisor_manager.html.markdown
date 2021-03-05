---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cisco_hypervisor_manager"
description: |-
  A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
---

# Data Source: intersight_hyperflex_cisco_hypervisor_manager
A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cisco_hypervisor_manager.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `identity`:(string) Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. 
* `nr_version`:(string) Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). 
 