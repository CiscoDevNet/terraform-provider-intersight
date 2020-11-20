
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_disk"
sidebar_current: "docs-intersight-data-source-storage-hitachi-disk"
description: |-
Disk entity associated with Hitachi storage array.
---

# Data Source: intersight_storage._hitachi_disk
Disk entity associated with Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `drive_type_code`:(string) Drive type code. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Disk name available in storage array. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `parity_group_id`:(string) Parity group number. When the drive does not belong to any parity group, an empty character string is output. 
* `part_number`:(string) Storage disk part number. 
* `protocol`:(string) Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe.* `Unknown` - Disk protocol is unknown.* `SAS` - Serial Attached SCSI protocol (SAS) used in disk.* `NVMe` - Non-volatile memory express (NVMe) protocol used in disk.* `SATA` - Serial Advanced Technology Attachment (SATA) used in disk. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `speed`:(int) Disk speed for read or write operation measured in rpm. 
* `status`:(string) Storage disk health status.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `type`:(string) Storage disk type - it can be SSD, HDD, NVRAM.* `Unknown` - Default unknown disk type.* `SSD` - Storage disk with Solid state disk.* `HDD` - Storage disk with Hard disk drive.* `NVRAM` - Storage disk with Non-volatile random-access memory type.* `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA).* `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.* `FC` - Storage disk with Fiber Channel.* `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.* `LUN` - Logical Unit Number refers to a logical disk.* `MSATA` - SATA disk in multi-disk carrier storage shelf.* `SAS` - Storage disk with serial attached SCSI.* `VMDISK` - Virtual machine Data Disk. 
* `type_detail`:(string) Drive type.* `N/A` - Not available.* `SAS` - SAS.* `SSD(MLC)` - SSD (MLC).* `SSD(FMC)` - SSD (FMC).* `SSD(FMD)` - SSD (FMD).* `SSD(SLC)` - SSD (SLC).* `SSD` - SSD.* `SSD(RI)` - SSD (RI). 
* `usage`:(string) Purpose for which the drive is used. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `version`:(string) Storage disk version number. 
