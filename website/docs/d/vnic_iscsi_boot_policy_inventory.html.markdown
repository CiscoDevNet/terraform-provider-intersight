---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_boot_policy_inventory"
description: |-
        Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.

---

# Data Source: intersight_vnic_iscsi_boot_policy_inventory
Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_iscsi_boot_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_targetvendor_name`:(string) Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `initiator_ip_source`:(string) Source Type of Initiator IP Address - Auto/Static/Pool.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `initiator_static_ip_v4_address`:(string) Static IP address provided for iSCSI Initiator. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_source_type`:(string) Source Type of Targets - Auto/Static.* `Static` - Type indicates that static target interface is assigned to iSCSI boot.* `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot. 
 
