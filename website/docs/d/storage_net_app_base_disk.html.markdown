---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_base_disk"
description: |-
        NetApp base disk is a storage array disk.

---

# Data Source: intersight_storage_net_app_base_disk
NetApp base disk is a storage array disk.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_base_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `base_disk_model`:(string) The NetApp base disk model. 
* `container_type`:(string) Supported container type for NetApp disk.* `Unknown` - Default container type is currently unknown.* `Aggregate` - Disk is used as a physical disk in an aggregate.* `Broken` - Disk is in a broken pool.* `Label Maintenance` - Disk is in online label maintenance list.* `Foreign` - Array LUN has been marked foreign.* `Maintenance` - Disk is in maintenance center.* `Mediator` - A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes.* `Shared` - Disk is partitioned or in a storage pool.* `Remote` - Disk belongs to a remote cluster.* `Spare` - The disk is a spare disk.* `Unassigned` - Disk ownership has not been assigned.* `Unsupported` - The disk is not supported. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `disk_bay`:(int) NetApp base disk shelf bay. 
* `disk_serial_number`:(string) NetApp base disk serial number. 
* `disk_shelf_id`:(string) NetApp base disk shelf id. 
* `disk_shelf_model`:(string) NetApp base disk shelf model. 
* `disk_shelf_name`:(string) NetApp base disk shelf name. 
* `disk_type`:(string) The type of the NetApp disk.* `Unknown` - Default unknown disk type.* `SSDNVM` - Solid state disk with Non-Volatile Memory Express protocol enabled.* `ATA` - Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself.* `FCAL` - For the FC-AL disk connection type, disk shelves are connected to the controller in a loop.* `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.* `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives.* `LUN` - Logical Unit Number refers to a logical disk.* `SAS` - Storage disk with serial attached SCSI.* `MSATA` - SATA disk in multi-disk carrier storage shelf.* `SSD` - Storage disk with Solid state disk.* `VMDISK` - Virtual machine Data Disk. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Disk name available in storage array. 
* `part_number`:(string) Storage disk part number. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `protocol`:(string) Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe.* `Unknown` - Disk protocol is unknown.* `SAS` - Serial Attached SCSI protocol (SAS) used in disk.* `NVMe` - Non-volatile memory express (NVMe) protocol used in disk.* `SATA` - Serial Advanced Technology Attachment (SATA) used in disk. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(int) Disk speed for read or write operation measured in rpm. 
* `state`:(string) Current state of the NetApp disk.* `Present` - Storage disk state type is present.* `Copy` - Storage disk state type is copy.* `Broken` - Storage disk state type is broken.* `Maintenance` - Storage disk state type is maintenance.* `Partner` - Storage disk state type is partner.* `Pending` - Storage disk state type is pending.* `Reconstructing` - Storage disk state type is reconstructing.* `Removed` - Storage disk state type is removed.* `Spare` - Storage disk state type is spare.* `Unfail` - Storage disk state type is unfail.* `Zeroing` - Storage disk state type is zeroing. 
* `status`:(string) Storage disk health status.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `type`:(string) Storage disk type - it can be SSD, HDD, NVRAM.* `Unknown` - Default unknown disk type.* `SSD` - Storage disk with Solid state disk.* `HDD` - Storage disk with Hard disk drive.* `NVRAM` - Storage disk with Non-volatile random-access memory type.* `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA).* `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.* `FC` - Storage disk with Fiber Channel.* `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.* `LUN` - Logical Unit Number refers to a logical disk.* `MSATA` - SATA disk in multi-disk carrier storage shelf.* `SAS` - Storage disk with serial attached SCSI.* `VMDISK` - Virtual machine Data Disk. 
* `uuid`:(string) Universally unique identifier of the NetApp Disk. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Storage disk version number. 
 
