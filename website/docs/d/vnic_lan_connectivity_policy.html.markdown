---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_lan_connectivity_policy"
description: |-
  A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.
---

# Data Source: intersight_vnic_lan_connectivity_policy
A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `iqn_allocation_type`:(string) Allocation Type of iSCSI Qualified Name - Static/Pool/None.* `None` - Type indicates that there is no IQN associated to an interface.* `Static` - Type represents that static IQN is associated to an interface.* `Pool` - Type indicates that IQN value is sourced from an associated pool. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `placement_mode`:(string) The mode used for placement of vNICs on network adapters. It can either be Auto or Custom.* `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.* `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. 
* `static_iqn_name`:(string) User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
