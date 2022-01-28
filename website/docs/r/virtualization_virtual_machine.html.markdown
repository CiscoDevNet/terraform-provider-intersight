---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_machine"
description: |-
  Depicts operations to control the life cycle of a virtual machine on a hypervisor.
---

# Resource: intersight_virtualization_virtual_machine
Depicts operations to control the life cycle of a virtual machine on a hypervisor.
## Usage Example
### Resource Creation

```hcl
resource "intersight_virtualization_virtual_machine" "virtualization_virtual_machine1" {
  name           = "virtualization_virtual_machine1"
  power_state    = "PowerOn"
  provision_type = "OVA"
  cpu            = 16
  memory         = 3755356
   host {
     object_type = "virtualization.BaseHost.relationship"
     moid        = var.virtualization_base_host_relationship
   }
   inventory {
     object_type = "virtualization.BaseVirtualMachine.relationship"
     moid        = var.virtualization_base_virtual_machine_relationship
   }
   registered_device {
     object_type = "asset.DeviceRegistration.relationship"
     moid        = var.DeviceRegistration_relationship
   }
   workflow_info {
     object_type = "workflow.WorkflowInfos.relationship"
     moid        = var.workflow_WorkflowInfo_relationship
   }
}

 variable "kubernetes_VirtualMachineInfrastructureProvider" {
   type = string
   description = "moid for kubernetes virtaul machine infrastructure provider"
 }

 variable "virtualization_base_host_relationship" {
   type = string
   description = "moid for kubernetes virtualization  base host relationship"
 }

 variable "virtualization_base_virtual_machine_relationship" {
   type = string
   description = "moid for  virtualization base virtualization machine relationship"
 }

 variable "DeviceRegistration_relationship" {
   type = string
   description = "moid for device registration relationship"
 }

 variable "workflow_WorkflowInfo_relationship" {
   type = string
   description = "moid for workflow workflowInfo relationship"
 }
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc).* `None` - A place holder for the default value.* `PowerState` - Power action is performed on the virtual machine.* `Migrate` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.* `Create` - The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation.* `Delete` - The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object. 
* `action_info`:(HashMap) -(ReadOnly) Details of an action performed on the virtual machine. Contains name of the action performed, status, failure reason message etc. 
This complex property has following sub-properties:
  + `failure_reason`:(string)(ReadOnly) Description of reason for failure. Derived from the workflow failure message. 
  + `name`:(string)(ReadOnly) Name of the Action performed on a resource like Virtual Machine, Disk etc. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string)(ReadOnly) Status of the Action like InProgress, Success, Failure etc.* `None` - A place holder for the default value.* `InProgress` - Previous action triggered on the resource is still running.* `Success` - Previous action triggered on the resource has completed successfully.* `Failure` - Previous action triggered on the resource has failed. 
* `affinity_selectors`:(Array)
This complex property has following sub-properties:
  + `name`:(string)(ReadOnly) Name of the meta property which identifies a specific resource. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string)(ReadOnly) Value of the meta property which identifies a specific resource. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `anti_affinity_selectors`:(Array)
This complex property has following sub-properties:
  + `name`:(string)(ReadOnly) Name of the meta property which identifies a specific resource. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string)(ReadOnly) Value of the meta property which identifies a specific resource. 
* `cloud_init_config`:(HashMap) - Cloud init configuration data for virtual machine. 
This complex property has following sub-properties:
  + `config_type`:(string) Virtual machine cloud init configuration type.* `` - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected.* `NoCloudSource` - Allows the user to provide user-data to the instance without running a network service.* `CloudConfigDrive` - Allows the user to provide user-data and network-data from cloud. 
  + `network_data`:(string) Network configuration data for a virtual machine. 
  + `network_data_base64_encoded`:(bool) Set to true, if the cloud init network data is in base64 format. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `user_data`:(string) User configuration data for a virtual machine such as adding user, group etc. 
  + `user_data_base64_encoded`:(bool) Set to true, if the cloud init user data is in base64 format. 
* `cluster`:(HashMap) - A reference to a virtualizationBaseCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cluster_esxi`:(string) Cluster where virtual machine is deployed. 
* `cpu`:(int) Number of vCPUs allocated to virtual machine. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `discovered`:(bool)(ReadOnly) Flag to indicate whether the configuration is created from inventory object. 
* `disk`:(Array)
This complex property has following sub-properties:
  + `bus`:(string) Disk bus name given for a virtual machine.* `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration.* `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives.* `scsi` - SCSI (Small Computer System Interface) bus used.. 
  + `name`:(string) Virtual machine network bridge name. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `order`:(int) Priority order of the disk. 
  + `type`:(string) Disk type hdd or cdrom for a virtual machine.* `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image.* `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image. 
  + `virtual_disk`:(HashMap) - Virtual disk configuration. 
This complex property has following sub-properties:
    + `capacity`:(string) Disk capacity to be provided with units example - 10Gi. 
    + `mode`:(string) File mode of the disk, example - Filesystem, Block.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk.* `` - Disk mode is either unknown or not supported. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `source_certs`:(string) Base64 encoded CA certificates of the https source to check against. 
    + `source_disk_to_clone`:(string) Source disk name from where the clone is done. 
    + `source_file_path`:(string) Disk image source for the virtual machine. 
  + `virtual_disk_reference`:(string) Name of the existing virtual disk to be attached to the Virtual Machine. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `force_delete`:(bool) Normally any virtual machine that is still powered on cannot be deleted. The expected sequence from a user is to first power off the virtual machine and then invoke the delete operation. However, in special circumstances, the owner of the virtual machine may know very well that the virtual machine is no longer needed and just wants to dispose it off. In such situations a delete operation of a virtual machine object is accepted only when this forceDelete attribute is set to true. Under normal circumstances (forceDelete is false), delete operation first confirms that the virtual machine is powered off and then proceeds to delete the virtual machine. 
* `guest_os`:(string) Guest operating system running on virtual machine.* `linux` - A Linux operating system.* `windows` - A Windows operating system. 
* `host`:(HashMap) - A reference to a virtualizationBaseHost resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `host_esxi`:(string) Host where virtual machine is deployed. 
* `hypervisor_type`:(string)(ReadOnly) Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `interfaces`:(Array)
This complex property has following sub-properties:
  + `adaptor_type`:(string) Virtual machine network adaptor type.* `Unknown` - The type of the network adaptor type is unknown.* `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC.* `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support.* `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.* `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features.* `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later.* `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets.* `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length.* `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance.* `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \ virtual\  devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \ emulated\  devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware.* `` - Default network adaptor type supported by the hypervisor. 
  + `bridge`:(string) Virtual machine network bridge name. 
  + `connect_at_power_on`:(bool) Connect the adaptor at virtual machine power on. 
  + `direct_path_io`:(bool) Enable the direct path I/O. 
  + `ip_forwarding_enabled`:(bool) Set to true, if IP forwarding is enabled on the NIC. 
  + `ipv6_address`:(bool) Set to true, if IPv6 address should be allocated for the NIC. 
  + `mac_address`:(string) Virtual machine network mac address. 
  + `name`:(string) Name of the network interface. This may be different from guest operating system assigned. 
  + `network_id`:(string) Identity of the network to which this network interface belongs. 
  + `nic_id`:(string) Identity of the network interface. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `order`:(int) Order of the NIC attachment to the VM. 
  + `private_ip_allocation_mode`:(string) Allocation mode for NIC addresses e.g. DHCP or static.* `DHCP` - Dynamic IP address allocation using DHCP protocol.* `STATIC_IP` - Assign fixed / static IPs to resources for use.* `IPAM_CALLOUT` - Use callout scripts to query cloud IP allocation tools to assign network parameters.* `PREALLOCATE_IP` - Allows the cloud infrastructure IP allocation to be dynamically provided before the server boots up. 
  + `public_ip_allocate`:(bool) Set to true, if public IP should be allocated for the NIC. 
  + `security_groups`:
                (Array of schema.TypeString) -
  + `static_ip_address`:(Array)
This complex property has following sub-properties:
    + `gateway_ip`:(string) IP address of the device on network which forwards local traffic to other networks. 
    + `ip_address`:(string) An IP address is a 32-bit number. It uniquely identifies a host in given network. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `subnet_mask`:(string) A 32 bit number which helps to identify the host and rest of the network. 
  + `subnet_id`:(string) Subnet identifier for the NIC. 
* `inventory`:(HashMap) -(ReadOnly) A reference to a virtualizationBaseVirtualMachine resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `labels`:(Array)
This complex property has following sub-properties:
  + `name`:(string)(ReadOnly) Name of the meta property which identifies a specific resource. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string)(ReadOnly) Value of the meta property which identifies a specific resource. 
* `memory`:(int) Virtual machine memory in mebi bytes (one mebibyte, 1MiB, is 1048576 bytes, and 1KiB is 1024 bytes). Input must be a whole number and scientific notation is not acceptable. For example, enter 1730 and not 1.73e03. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Virtual machine name that is unique. Hypervisors enforce platform specific limits and character sets. The name length limit, both min and max, vary among hypervisors. Therefore, the basic limits are set here and proper enforcement is done elsewhere. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `power_state`:(string) Expected power state of virtual machine (PowerOn, PowerOff, Restart).* `PowerOff` - The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected.* `PowerOn` - The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected.* `Suspend` - The virtual machine will be put into  a suspended state.* `ShutDownGuestOS` - The guest operating system is shut down gracefully.* `RestartGuestOS` - It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts.* `Reset` - Resets the virtual machine abruptly, with no consideration for work in progress.* `Restart` - The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up.* `Unknown` - Power state of the entity is unknown. 
* `provision_type`:(string) Identifies the provision type to create a new virtual machine.* `OVA` - Deploy virtual machine using OVA/F file.* `Template` - Provision virtual machine using a template file.* `Discovered` - A virtual machine was 'discovered' and not created from Intersight. No provisioning information is available. 
* `registered_device`:(HashMap) - A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vm_config`:(HashMap) - Virtual machine configuration to provision. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `workflow_info`:(HashMap) -(ReadOnly) A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_virtualization_virtual_machine` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_virtualization_virtual_machine.example 1234567890987654321abcde
``` 