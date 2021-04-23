---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_disk"
description: |-
  Disk entity associated with Pure FlashArray.
---

# Data Source: intersight_storage_pure_disk
Disk entity associated with Pure FlashArray.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
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
* `status`:(string) Storage disk health status.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `type`:(string) Storage disk type - it can be SSD, HDD, NVRAM.* `Unknown` - Default unknown disk type.* `SSD` - Storage disk with Solid state disk.* `HDD` - Storage disk with Hard disk drive.* `NVRAM` - Storage disk with Non-volatile random-access memory type.* `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA).* `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.* `FC` - Storage disk with Fiber Channel.* `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.* `LUN` - Logical Unit Number refers to a logical disk.* `MSATA` - SATA disk in multi-disk carrier storage shelf.* `SAS` - Storage disk with serial attached SCSI.* `VMDISK` - Virtual machine Data Disk. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Storage disk version number. 
 