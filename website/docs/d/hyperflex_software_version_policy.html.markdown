---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_software_version_policy"
description: |-
  A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
---

# Data Source: intersight_hyperflex_software_version_policy
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_software_version_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `hxdp_version`:(string) Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster. 
* `hypervisor_version`:(string) Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `server_firmware_version`:(string) Desired server firmware version to apply on the HyperFlex Cluster. 
 