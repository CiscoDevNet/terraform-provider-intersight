---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_virtual_machine"
description: |-
  The Virtual machine that runs on a Hyperflex Application platform compute host.
---

# Data Source: intersight_hyperflex_hxap_virtual_machine
The Virtual machine that runs on a Hyperflex Application platform compute host.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_virtual_machine.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `age`:(string) Denotes age or life time of the VM in nano seconds. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `failure_reason`:(string) Reason of the failure when VM is in failed state. 
* `hypervisor_type`:(string) Type of hypervisor where the virtual machine is hosted for example ESXi.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the virtual machine. 
* `network_count`:(int) Number network interfaces associated with the virtual machine. 
* `power_state`:(string) Power state of the virtual machine.* `Unknown` - The entity's power state is unknown.* `PoweredOn` - The entity is powered on.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state.* `` - The entity's power state is not available. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Denotes the VM start timestamp. 
* `status`:(string) Status of virtual machine.* `Unknown` - Virtual machine state is not available.* `Running` - Virtual machine is running normally.* `Stopped` - Virtual machine has been stopped.* `WaitForLaunch` - Virtual machine is wating to be launched.* `Paused` - Virtual machine is currently paused.* `ImportInProgress` - Virtual machine image is being imported into the platform.* `ImportFailed` - Virtual machine image import operation failed.* `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress.* `DiskCloneFailed` - Disk clone operation for the virtual machine failed.* `Processing` - Virtual machine is being created.* `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.* `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation.* `` - Virtual machine status is not available. 
* `uuid`:(string) The uuid of this virtual machine. The uuid is internally generated and not user specified. 
 