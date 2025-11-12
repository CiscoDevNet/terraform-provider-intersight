---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_ahv_vm_nic"
description: |-
        A NIC associated with an AHV VM.

---

# Data Source: intersight_hci_ahv_vm_nic
A NIC associated with an AHV VM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_ahv_vm_nic.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_connected`:(bool) Indicates if the NIC is connected. 
* `mac_address`:(string) The MAC address of the NIC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model name of emulated NIC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `network_function_nic_type`:(string) The type of this network function NIC. 
* `nic_ext_id`:(string) The unique identifier of the NIC. 
* `nic_type`:(string) The type of the NIC. Possible values are 'NORMAL_NIC', 'DIRECT_NIC', 'NETWORK_FUNCTION_NIC', 'SPAN_DESTINATION_NIC'. 
* `num_queues`:(int) The number of Tx/Rx queue pairs for this NIC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `should_allow_unknown_macs`:(bool) Indicates whether an unknown unicast traffic is forwarded to this NIC or not, only for the NICs on the overlay subnets. 
* `vlan_mode`:(string) By default, all the virtual NICs are created in ACCESS mode, which permits only one VLAN per virtual network. TRUNKED mode allows multiple VLANs on a single VM NIC for network-aware user VMs. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
 
