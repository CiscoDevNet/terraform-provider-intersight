---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_policy"
description: |-
  A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.
---

# Data Source: intersight_fabric_eth_network_policy
A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_eth_network_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 