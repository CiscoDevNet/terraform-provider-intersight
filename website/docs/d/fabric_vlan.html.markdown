---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vlan"
description: |-
  Configuration object for Virtual LAN.
---

# Data Source: intersight_fabric_vlan
Configuration object for Virtual LAN.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `auto_allow_on_uplinks`:(bool) Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI. 
* `is_native`:(bool) Used to define whether this VLAN is to be classified as 'native' for traffic in this FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The 'name' used to identify this VLAN. 
* `vlan_id`:(int) The identifier for this Virtual LAN. 
