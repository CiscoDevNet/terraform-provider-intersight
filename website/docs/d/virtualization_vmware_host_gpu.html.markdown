---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_host_gpu"
description: |-
        Common attributes of a GPU device on a VMware host.

---

# Data Source: intersight_virtualization_vmware_host_gpu
Common attributes of a GPU device on a VMware host.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_host_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bus`:(int) The bus ID of this PCI device. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(int) The device ID of this PCI device. 
* `device_name`:(string) The device name of this PCI device. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `function`:(int) The function ID of this PCI device. 
* `identity`:(string) The internally generated identity of this PCI device. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `passthrough_active`:(bool) Indicates if passthrough is active for this PCI device (meaning enabled + rebooted). 
* `passthrough_enabled`:(bool) Indicates if passthrough feature is enabled for this PCI device.PCI passthrough feature enables you to access and manage hardware devices from a virtual machine.When PCI passthrough is configured, the PCI devices function as if they were physically attached to the guest operating system. 
* `pci_class_id`:(int) The class ID of this PCI device. 
* `pci_id`:(string) The ID of this PCI device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot`:(int) The slot ID of this PCI device. 
* `sub_device_id`:(int) The sub device ID of this PCI device. 
* `sub_vendor_id`:(int) The sub vendor ID of this PCI device. 
* `vendor_id`:(int) The vendor ID of this PCI device. 
* `vendor_name`:(string) The vendor name of this PCI device. 
 
