---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_machine_gpu"
description: |-
        Common attributes of virtual GPU device on a VMware virtual machine.

---

# Data Source: intersight_virtualization_vmware_virtual_machine_gpu
Common attributes of virtual GPU device on a VMware virtual machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_virtual_machine_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backing_pci_id`:(string) The backing physical host PCI device Id for this device. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(int) The internally assigned key of this virtual GPU device. This entity is not manipulated by users. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of this virtual machine PCI device. 
* `passthrough`:(bool) Indicates if this virtual machine PCI device is enabled via passthrough from the host. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vm_identity`:(string) Identity of the virtual machine where the virtual gpu is created. 
 
