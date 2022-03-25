---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_aws_virtual_machine"
description: |-
        Configuration details of the virtual machine.

---

# Data Source: intersight_cloud_aws_virtual_machine
Configuration details of the virtual machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_aws_virtual_machine.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `boot_time`:(string) Time when this VM booted up. 
* `cpu_utilization`:(float) Average CPU utilization percentage derived as a ratio of CPU used to CPU allocated. The value is calculated whenever inventory is performed. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hypervisor_type`:(string) Type of hypervisor where the virtual machine is hosted for example ESXi.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). 
* `memory_utilization`:(float) Average memory utilization percentage derived as a ratio of memory used to available memory. The value is calculated whenever inventory is performed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the virtual machine. 
* `power_state`:(string) Power state of the virtual machine.* `Unknown` - The entity's power state is unknown.* `PoweringOn` - The entity is powering on.* `PoweredOn` - The entity is powered on.* `PoweringOff` - The entity is powering off.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state.* `Rebooting` - The entity reboot is in progress.* `` - The entity's power state is not available. 
* `private_dns`:(string) The private DNS hostname name assigned to the instance. 
* `nr_provider`:(string) Cloud platform, where the virtual machine is launched.* `Unknown` - Cloud provider is not known.* `VMwarevSphere` - Cloud provider named VMware vSphere.* `AmazonWebServices` - Cloud provider named Amazon Web Services.* `MicrosoftAzure` - Cloud provider named Microsoft Azure.* `GoogleCloudPlatform` - Cloud provider named Google Cloud Platform.* `CiscoIntersightWorkloadEngine` - Cloud provider named Cisco Intersight Workload Engine. 
* `public_dns`:(string) The public DNS name assigned to the instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The current state of the virtual machine. For example, starting, stopped, etc.* `None` - A place holder for the default value.* `Creating` - Virtual machine creation is in progress.* `Pending` - The virtual machine is preparing to enter the started state.* `Starting` - The virtual machine is starting.* `Started` - The virtual machine is running and ready for use.* `Stopping` - The virtual machine is preparing to be stopped.* `Stopped` - The virtual machine is shut down and cannot be used. The virtual machine can be started again at any time.* `Pausing` - The virtual machine is preparing to be paused.* `Paused` - The virtual machine enters into paused state due to low free disk space.* `Suspending` - The virtual machine is preparing to be suspended.* `Suspended` - Virtual machine is in sleep mode.When a virtual machine is suspended, the current state of theoperating system, and applications is saved, and the virtual machine put into a suspended mode.* `Deleting` - The virtual machine is preparing to be terminated.* `Terminated` - The virtual machine has been permanently deleted and cannot be started.* `Rebooting` - The virtual machine reboot is in progress.* `Error` - The deployment of virtual machine is failed.* `Warning` - The virtual machine is in warning state. 
* `tenancy`:(string) How virtual machines are distributed across physical hardware and affects pricing. 
* `termination_time`:(string) Time when this virtualmachine is terminated. 
* `uuid`:(string) The uuid of this virtual machine. The uuid is internally generated and not user specified. 
* `vm_creation_time`:(string) Time when this virtualmachine is created. 
 
