---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_network_interface"
description: |-
  Details of VMware virtual network interface.
---

# Data Source: intersight_virtualization_vmware_virtual_network_interface
Details of VMware virtual network interface.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_virtual_network_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_type`:(string) Type of virtual ethernet adapter for virtual network interface. 
* `connect_at_power_on`:(bool) Connect or not to connect the device when the virtual machine starts. 
* `connected`:(bool) Device is currently connected or not. Valid only while the virtual machine is running. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(int) The internally assigned key of this virtual network interface. This entity is not manipulated by users. 
* `mac_address`:(string) MAC address assigned to virtual network interface. 
* `mac_address_type`:(string) MAC address type for the mac address assigned to virtual network interface.* `manual` - Statically assigned MAC address.* `generated` - Automatically generated MAC address.* `assigned` - MAC address assigned by VCenter to the virtual network interface card. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual network interface created on a virtual machine. 
* `network_type`:(string) Type of network for virtual network interface. It can be either standard or distributed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vm_identity`:(string) Identity of the virtual machine where the virtual network interface is created. 
 