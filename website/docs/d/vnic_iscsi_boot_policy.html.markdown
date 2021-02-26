---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_boot_policy"
description: |-
  Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
---

# Data Source: intersight_vnic_iscsi_boot_policy
Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `auto_targetvendor_name`:(string) Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters. 
* `description`:(string) Description of the policy. 
* `initiator_ip_source`:(string) Source Type of Initiator IP Address - Auto/Static/Pool.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `initiator_static_ip_v4_address`:(string) Static IP address provided for iSCSI Initiator. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `target_source_type`:(string) Source Type of Targets - Auto/Static.* `Static` - Type indicates that static target interface is assigned to iSCSI boot.* `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot. 
