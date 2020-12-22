---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_port_channel"
description: |-
  Model contains the details of the ethernet port-channels configured on the FI.
---

# Data Source: intersight_ether_port_channel
Model contains the details of the ethernet port-channels configured on the FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_vlan`:(string) Access VLANs for this port-channel, on this FI. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port-channel. 
* `allowed_vlans`:(string) Allowed VLANs on this port-channel, on this FI. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `mode`:(string) Operating mode of this port-channel. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `native_vlan`:(string) Native VLAN for this port-channel, on this FI. 
* `oper_speed`:(string) Operational speed of this port-channel. 
* `oper_state`:(string) Operational state of this port-channel. 
* `oper_state_qual`:(string) Reason for this port-channel's Operational state. 
* `port_channel_id`:(int) Unique identifier for this port-channel on the FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) This port-channel's configured role (uplink, server, etc.). 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
