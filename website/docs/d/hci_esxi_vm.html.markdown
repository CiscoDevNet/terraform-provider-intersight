---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_esxi_vm"
description: |-
        An ESXi VM reported by Prism Central.

---

# Data Source: intersight_hci_esxi_vm
An ESXi VM reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_esxi_vm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The unique identifier of the cluster which owns this VM. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of the VM. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `guest_os_name`:(string) The guest OS name of the VM. 
* `hypervisor_type`:(string) The hypervisor type of the given VM. It could be AHV, ESX etc. 
* `memory_size_bytes`:(int) The memory size in bytes of the VM. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the VM reported by the Prism Central. 
* `node_ext_id`:(string) The unique identifier of the node. 
* `num_cores_per_socket`:(int) The number of cores per socket of the VM. 
* `num_cpu_cores`:(int) The number of CPU cores of the VM. 
* `power_state`:(string) The power state of the VM. The possible values are ON, OFF, SUSPENDED (ESXi), PAUSED (AHV), UNDETERMINED.* `UNDETERMINED` - The VM power state is currently unknown.* `OFF` - The VM's power state is powered-off.* `ON` - The VM's power state is powered-on.* `PAUSED` - The VM's power state is paused, applicable only to AHV VM.* `SUSPENDED` - The VM's power state is suspended, applicable only to ESXi VM. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `virtual_hardware_version`:(int) The virtual hardware version of the VM. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
 
