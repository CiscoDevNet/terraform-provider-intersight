---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_machine"
description: |-
  Depicts operations to control the life cycle of a virtual machine on a hypervisor.
---

# Data Source: intersight_virtualization_virtual_machine
Depicts operations to control the life cycle of a virtual machine on a hypervisor.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_virtual_machine.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc).* `None` - A place holder for the default value.* `PowerState` - Power action is performed on the virtual machine.* `Migrate` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.* `Create` - The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation.* `Delete` - The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object. 
* `cluster_esxi`:(string) Cluster where virtual machine is deployed. 
* `cpu`:(int) Number of vCPUs allocated to virtual machine. 
* `create_time`:(string) The time when this managed object was created. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `force_delete`:(bool) Normally any virtual machine that is still powered on cannot be deleted. The expected sequence from a user is to first power off the virtual machine and then invoke the delete operation. However, in special circumstances, the owner of the virtual machine may know very well that the virtual machine is no longer needed and just wants to dispose it off. In such situations a delete operation of a virtual machine object is accepted only when this forceDelete attribute is set to true. Under normal circumstances (forceDelete is false), delete operation first confirms that the virtual machine is powered off and then proceeds to delete the virtual machine. 
* `guest_os`:(string) Guest operating system running on virtual machine.* `linux` - A Linux operating system.* `windows` - A Windows operating system. 
* `host_esxi`:(string) Host where virtual machine is deployed. 
* `hypervisor_type`:(string) Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `memory`:(int) Virtual machine memory in mebi bytes (one mebibyte, 1MiB, is 1048576 bytes, and 1KiB is 1024 bytes). Input must be a whole number and scientific notation is not acceptable. For example, enter 1730 and not 1.73e03. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Virtual machine name that is unique. Hypervisors enforce platform specific limits and character sets. The name length limit, both min and max, vary among hypervisors. Therefore, the basic limits are set here and proper enforcement is done elsewhere. 
* `power_state`:(string) Expected power state of virtual machine (PowerOn, PowerOff, Restart).* `PowerOff` - The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected.* `PowerOn` - The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected.* `Suspend` - The virtual machine will be put into  a suspended state.* `ShutDownGuestOS` - The guest operating system is shut down gracefully.* `RestartGuestOS` - It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts.* `Reset` - Resets the virtual machine abruptly, with no consideration for work in progress.* `Restart` - The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up.* `Unknown` - Power state of the entity is unknown. 
* `provision_type`:(string) Identifies the provision type to create a new virtual machine.* `OVA` - Deploy virtual machine using OVA/F file.* `Template` - Provision virtual machine using a template file.* `Discovered` - A virtual machine was 'discovered' and not created from Intersight. No provisioning information is available. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 