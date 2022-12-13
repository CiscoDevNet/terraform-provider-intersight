---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_drive"
description: |-
        A Hyperflex drive entity attached to a node in a
        Hyperflex cluster.

---

# Data Source: intersight_hyperflex_drive
A Hyperflex drive entity attached to a node in a
Hyperflex cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_drive.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(int) Provisioned capacity of the drive in bytes. 
* `create_time`:(string) The time when this managed object was created. 
* `disk_use_state`:(string) Disk inventory state. Should be one of values defined in enum.* `UNKNOWN` - The use state of the disk is unknown.* `USED` - The use state of the disk is used.* `NOTUSED` - The use state of the disk is unused.* `EMPTY` - The use state of the disk is empty. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name`:(string) Host Name. 
* `host_uuid`:(string) The unique identifier of the drive's host. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model_number`:(string) The model number of the disk or SSD. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_uuid`:(string) The unique identifier of the Hyperflex Node to which the disk is attached. 
* `path`:(string) Device path for the drive. 
* `protocol`:(string) Disk Protocol - SATA, NVME, SAS.* `Unknown` - Disk protocol is unknown.* `SAS` - Serial Attached SCSI protocol (SAS) used in disk.* `NVMe` - Non-volatile memory express (NVMe) protocol used in disk.* `SATA` - Serial Advanced Technology Attachment (SATA) used in disk. 
* `serial_number`:(string) Serial number of the Hyperflex drive. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_number`:(string) The SCSI slot number of the drive. 
* `status`:(string) Disk inventory state as determined by storfs inventory module.Should be one of values defined in enum.* `UNKNOWN` - The state of the disk is unknown.* `CLAIMED` - The state of the disk is claimed by storfs and has a valid storfs format.* `AVAILABLE` - The disk is available but not claimed by storfs.* `IGNORED` - The disk ash been ignored by storfs.* `BLACKLISTED` - The deprecated value for 'Blocked'. It is included to maintain backwards compatibility with clusters running a HyperFlex Data Platform version older than 5.0(1a).* `SECUREERASED` - The disk has been secure erased.* `BLOCKED` - The disk has been blocked by storfs. 
* `type`:(string) Type of disk - UNKNOWN, HDD, SSD.* `Unknown` - Default unknown disk type.* `SSD` - Storage disk with Solid state disk.* `HDD` - Storage disk with Hard disk drive.* `NVRAM` - Storage disk with Non-volatile random-access memory type.* `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA).* `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.* `FC` - Storage disk with Fiber Channel.* `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.* `LUN` - Logical Unit Number refers to a logical disk.* `MSATA` - SATA disk in multi-disk carrier storage shelf.* `SAS` - Storage disk with serial attached SCSI.* `VMDISK` - Virtual machine Data Disk. 
* `usage`:(string) Specifies what the disk is used for. Should be one ofvalues defined in enum.* `UNKNOWN` - The usage of the disk is unknown.* `PERSISTENCE` - The usage of the disk is for generic data storage.* `SYSTEM` - The usage of the disk is related to system storage.* `CACHING` - The usage of the disk is for caching. 
* `used_capacity`:(int) Used Capacity of the drive in Bytes. 
* `uuid`:(string) The unique identifier of the Hyperflex drive. 
* `nr_version`:(string) The firmware version of the Hyperflex drive. 
 
