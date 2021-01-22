---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_ethernet_port"
description: |-
  Ethernet port is a port on a node in a storage array.
---

# Data Source: intersight_storage_net_app_ethernet_port
Ethernet port is a port on a node in a storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enabled`:(string) Status of Port to determine if its enabled or not. 
* `mac_address`:(string) Macaddress  of the port available in storage array. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(string) Maximum transmission unit of the physical port available in storage array. 
* `name`:(string) Name of the port available in storage array. 
* `speed`:(int) Operational speed of port measured. 
* `state`:(string) State of the port available in storage array.* `down` - An inactive port is listed as Down.* `up` - An active port is listed as Up.* `present` - An active port is listed as present. 
* `type`:(string) Type of the port available in storage array.* `LAG` - Storage port of type lag.* `physical` - LIFs can be configured directly on physical ports.* `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port. 
* `uuid`:(string) UUID of physical port. 
