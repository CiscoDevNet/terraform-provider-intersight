---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_network_policy"
description: |-
  A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.
---

# Data Source: intersight_vnic_fc_network_policy
A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_fc_network_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 