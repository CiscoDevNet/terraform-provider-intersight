
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_volume"
sidebar_current: "docs-intersight-data-source-storage-hitachi-volume"
description: |-
A volume entity in Hitachi storage array.
---

# Data Source: intersight_storage._hitachi_volume
A volume entity in Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `clpr_id`:(int) CLPR (Cache Logical Partition) number of this volume. 
* `data_reduction_mode`:(string) Setting of the capacity saving function (dedupe and compression).* `N/A` - Not available.* `compression` - The capacity saving function (compression) is enabled.* `compression_deduplication` - The capacity saving function (compression and deduplication) is enabled.* `disabled` - The capacity saving function (compression and deduplication) is disabled. 
* `data_reduction_status`:(string) Status of the capacity saving function.* `N/A` - Not available.* `ENABLED` - The capacity saving function is enabled.* `DISABLED` - The capacity saving function is disabled.* `ENABLING` - The capacity saving function is being enabled.* `REHYDRATING` - The capacity saving function is being disabled.* `DELETING` - The volumes for which the capacity saving function is enabled are being deleted.* `FAILED` - An attempt to enable the capacity saving function failed. 
* `description`:(string) Short description about the volume. 
* `drive_type`:(string) Code indicating the drive type of the drive belonging to the volume. 
* `emulation_type`:(string) The volume emulation type or the volume status information.* `N/A` - Not available.* `NOT DEFINED` - The volume is not implemented.* `DEFINING` - The volume is being created.* `REMOVING` - The volume is being removed.* `OPEN-V` - To be provided by Hitachi. 
* `is_full_allocation_enabled`:(bool) Whether pages are reserved by the FullAllocation functionality. 
* `label`:(string) Label of the volume, as configured in the storage array. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `num_of_paths`:(int) Number of paths set for the volume. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pool_id`:(string) ID of the pool with which the volume is associated. 
* `raid_level`:(string) RAID level for the volume.* `N/A` - Not available.* `RAID1` - RAID1.* `RAID5` - RAID5.* `RAID6` - RAID6. 
* `raid_type`:(string) RAID type drive configuration. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `status`:(string) Status information of the volume.* `N/A` - Not available.* `NML` - The volume is in normal status.* `BLK` - The volume is blocked.* `BSY` - The volume status is being changed.* `Unknown` - The volume status is unknown (not supported). 
