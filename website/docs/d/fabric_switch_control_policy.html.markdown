---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_control_policy"
description: |-
  A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
---

# Data Source: intersight_fabric_switch_control_policy
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_control_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `vlan_port_optimization_enabled`:(bool) To enable or disable the VLAN port count optimization. 
 