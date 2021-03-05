---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_appliance_pc_role"
description: |-
  Configuration object sent by user to create an appliance port channel.
---

# Data Source: intersight_fabric_appliance_pc_role
Configuration object sent by user to create an appliance port channel.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_appliance_pc_role.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port channel.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps. 
* `mode`:(string) Port mode to be set on the appliance port-channel.* `trunk` - Trunk Mode Switch Port Type.* `access` - Access Mode Switch Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `priority`:(string) The 'name' of the System QoS Class.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
 