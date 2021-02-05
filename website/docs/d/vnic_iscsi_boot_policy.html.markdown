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
* `initiator_ip_source`:(string) Source Type of Initiator IP Address - Auto/Static/Pool.* `None` - Type defines that property is not applicable for an interface.* `Auto` - The system selects an interface automatically - DHCP.* `Static` - Type represents that static information or properties are associated to an interface.* `Pool` - Type defines that property value will be fetched from an associated pool. 
* `initiator_static_ip_v4_address`:(string) Static IP address provided for iSCSI Initiator. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `target_source_type`:(string) Source Type of Targets - Auto/Static.* `None` - Type defines that property is not applicable for an interface.* `Auto` - The system selects an interface automatically - DHCP.* `Static` - Type represents that static information or properties are associated to an interface.* `Pool` - Type defines that property value will be fetched from an associated pool. 
