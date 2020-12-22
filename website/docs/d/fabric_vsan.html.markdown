---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vsan"
description: |-
  Configuration object sent by user to create VSAN configurations.
---

# Data Source: intersight_fabric_vsan
Configuration object sent by user to create VSAN configurations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `default_zoning`:(string) Enables or Disables the default zoning state.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `fc_zone_sharing_mode`:(string) Logical grouping mode for fc ports. 
* `fcoe_vlan`:(int) FCOE Vlan associated to the VSAN configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User given name for the VSAN configuration. 
* `vsan_id`:(int) Virtual San Identifier in the switch. 
