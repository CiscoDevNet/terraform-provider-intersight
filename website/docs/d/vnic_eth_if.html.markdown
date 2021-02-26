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
* `iscsi_ip_v4_address_allocation_type`:(string) Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.* `None` - Type indicates that there is no IP associated to an vnic.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `iscsi_ipv4_address`:(string) IP address associated to the vNIC. 
* `mac_address`:(string) The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. 
* `mac_address_type`:(string) Type of allocation selected to assign a MAC address for the vnic.* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual ethernet interface. 
* `order`:(int) The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. 
* `standby_vif_id`:(int) The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. 
* `static_mac_address`:(string) The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use thefollowing MAC prefix 00:25:B5:xx:xx:xx. 
* `vif_id`:(int) The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. 
