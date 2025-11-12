---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_esxi_vm_nic"
description: |-
        A NIC associated with an ESXi VM.

---

# Data Source: intersight_hci_esxi_vm_nic
A NIC associated with an ESXi VM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_esxi_vm_nic.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_type`:(string) The adapter type of the NIC. Possible values are 'E1000', 'E1000E', 'VMXNET', 'VMXNET2', 'VMXNET3', 'PCNET32', 'SRIOV'. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_connected`:(bool) Indicates if the NIC is connected. 
* `mac_address`:(string) The MAC address of the NIC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nic_ext_id`:(string) The unique identifier of the NIC. 
* `portgroup_name`:(string) The name of the port group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
 
