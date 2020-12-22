---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_if"
description: |-
  Virtual Ethernet Interface.
---

# Data Source: intersight_vnic_eth_if
Virtual Ethernet Interface.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `failover_enabled`:(bool) Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover. 
* `mac_address`:(string) The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual ethernet interface. 
* `order`:(int) The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. 
* `standby_vif_id`:(int) The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. 
* `vif_id`:(int) The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. 
