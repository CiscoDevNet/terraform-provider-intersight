---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_kernel_network"
description: |-
        Details of VMware Kernel Network.

---

# Data Source: intersight_virtualization_vmware_kernel_network
Details of VMware Kernel Network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_kernel_network.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fault_tolerance_logging`:(bool) Indicates that fault tolerance logging is enabled on this kernel network. 
* `identity`:(string) The internally generated identity of network. This entity cannot manipulated by users. It aids in uniquely identifying the network object. For VMware, this is MOR (managed object reference). 
* `mac_address`:(string) Standard MAC address assigned to this kernel network. 
* `management`:(bool) Indicates that management traffic is enabled on this kernel network. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit configured on a kernel network. 
* `name`:(string) User-provided name to identify the portgroup. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subnet_mask`:(string) Subnet mask of the kernel network. 
* `tcp_ip_stack`:(string) Type of stack for the kernel network. It can be custom, default, vMotion or provisioning. 
* `vmotion`:(bool) Indicates that vmotion is enabled on this kernel network. 
* `vsan`:(bool) Indicates that vsan traffic is enabled on this kernel network. 
* `vsphere_provisioning`:(bool) Indicates that vsphere provisioning traffic is enabled on this kernel network. 
* `vsphere_replication`:(bool) Indicates that vsphere replication is enabled on this kernel network. 
* `vsphere_replication_nfc`:(bool) Indicates that vsphere replication nfc is enabled on this kernel network. 
 
