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
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_virtual_machine.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `age`:(string) Denotes age or life time of the VM in nano seconds. 
* `failure_reason`:(string) Reason of the failure when VM is in failed state. 
* `graphic_console_url`:(string) Graphical console URL of this VM. 
* `hypervisor_type`:(string) Type of hypervisor where the virtual machine is hosted for example ESXi.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the virtual machine. 
* `network_count`:(int) Number network interfaces associated with the virtual machine. 
* `power_state`:(string) Power state of the virtual machine.* `Unknown` - The entity's power state is unknown.* `PoweredOn` - The entity is powered on.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state. 
* `start_time`:(string) Denotes the VM start timestamp. 
* `status`:(string) Status of virtual machine.* `Unknown` - Virtual machine state is not available.* `Running` - Virtual machine is running normally.* `Stopped` - Virtual machine has been stopped.* `WaitForLaunch` - Virtual machine is wating to be launched.* `Paused` - Virtual machine is currently paused.* `ImportInProgress` - Virtual machine image is being imported into the platform.* `ImportFailed` - Virtual machine image import operation failed.* `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress.* `DiskCloneFailed` - Disk clone operation for the virtual machine failed.* `Processing` - Virtual machine is being created.* `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.* `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation. 
* `up_time`:(string) Denotes how long this VM has been running in nano seconds. 
* `uuid`:(string) The uuid of this virtual machine. The uuid is internally generated and not user specified. 
 