---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_disk"
description: |-
  Depicts disk configuration used to create a virtual disk on a hypervisor datastore.
---

# Data Source: intersight_virtualization_vmware_virtual_disk
Depicts disk configuration used to create a virtual disk on a hypervisor datastore.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_virtual_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `compatibility_mode`:(string) Compatibility mode of the raw disk mapping (RDM).* `notApplicable` - Value specified for any disk which is not of raw device mapping type.* `physicalMode` - A disk device backed by a physical compatibility mode raw disk mapping cannot use disk modes, and commands are passed straight through to the LUN indicated by the raw disk mapping.* `virtualMode` - A disk device backed by a virtual compatibility mode raw disk mapping can use disk modes. 
* `controller_key`:(int) Key of the controller on which the disk is created. 
* `create_time`:(string) The time when this managed object was created. 
* `device_name`:(string) Host-specific device the LUN is being accessed through. If the target LUN is not available on the host then it is empty. For example, this could happen if it has accidentally been masked out. 
* `disk_mode`:(string) Persistence mode of the virtual disk. For RDM disks, it is only supported when the raw disk mapping is using virtual compatibility mode.* `persistent` - Changes are immediately and permanently written to the virtual disk.* `independent_persistent` - Changes are immediately and permanently written to the virtual disk and not affected by snapshots.* `independent_nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off and not affected by snapshots.* `nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off.* `undoable` - Changes to virtual disk are made to a redo log and has the option to commit or undo.* `append` - Changes are appended to the redo log and can be revoked by removing the undo log. 
* `disk_type`:(string) Specifies whether the virtual disk is a RDM or a flat disk.* `flatDisk` - Specifies that it is a flat disk.* `rdmDisk` - Specifies that it is a raw device mapping disk. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(int) The internally assigned key of this disk. This entity is not manipulated by users. 
* `limit`:(int) The utilization of a virtual machine will not exceed this limit, even if there are available resources. Used to ensure a consistent performance of virtual machines independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). The unit is number of I/O per second. 
* `lun_uuid`:(string) Unique identifier of the LUN accessed by the raw disk mapping (RDM). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `serial`:(string) Serial ID of the storage device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sharing`:(string) Sharing mode of the virtual disk.* `sharingNone` - The virtual disk is not shared.* `sharingMultiWriter` - The virtual disk is shared between multiple virtual machines. 
* `size`:(int) Disk size represented in bytes. 
* `storage_allocation_type`:(string) Allocation type for the virtual disk.* `notApplicable` - Value specified for a disk for which storage allocation type is not applicable.* `thin` - A thin provisioned disk consumes only the space that it needs for its initial operrations, and grows with time according to demand. It is the fastest method to create a virtual disk because it creates a disk with just the header information.* `lazyZeroedThick` - A lazy zeroed thick disk has all space allocated at the time of its creation. Data remaining on the physical device is not erased during creation, but is zeroed out on demand later on first write from the virtual machine.* `eagerZeroedThick` - An eager zeroed thick disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. Such disks may take longer time during creation compared to other disk formats. 
* `unit_number`:(int) Unit number of the disk on its controller. 
* `uuid`:(string) UUID assigned by vCenter to every disk. 
* `vdisk_id`:(string) Identity of the virtual disk object as the first class entity. 
* `vendor`:(string) Vendor of the storage device. 
* `virtual_disk_path`:(string) Path of the virtual disk. 
* `vm_identity`:(string) Identity of the virtual machine where the virtual disk is created. 
 