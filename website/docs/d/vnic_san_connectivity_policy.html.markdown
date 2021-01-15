---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_san_connectivity_policy"
description: |-
  SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
---

# Data Source: intersight_vnic_san_connectivity_policy
SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `placement_mode`:(string) The mode used for placement of vNICs on network adapters. It can either be Auto or Custom.* `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.* `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. 
* `static_wwnn_address`:(string) The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF.To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
* `wwnn_address_type`:(string) Type of allocation selected to assign a WWNN address for the server node.* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface. 
