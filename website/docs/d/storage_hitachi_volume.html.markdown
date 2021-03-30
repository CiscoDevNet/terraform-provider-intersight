---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_volume"
description: |-
  A volume entity in Hitachi storage array.
---

# Data Source: intersight_storage_hitachi_volume
A volume entity in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_volume.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `clpr_id`:(int) CLPR (Cache Logical Partition) number of this volume. 
* `data_reduction_mode`:(string) Setting of the capacity saving function (dedupe and compression).* `N/A` - The capacity saving function is not available.* `Compression` - The capacity saving function (compression) is enabled.* `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled.* `Disabled` - The capacity saving function (compression and deduplication) is disabled. 
* `data_reduction_status`:(string) Status of the capacity saving function.* `N/A` - The capacity saving function is not available.* `Enabled` - The capacity saving function is enabled.* `Disabled` - The capacity saving function is disabled.* `Enabling` - The capacity saving function is being enabled.* `Rehydrating` - The capacity saving function is being disabled.* `Deleting` - The volumes for which the capacity saving function is enabled are being deleted.* `Failed` - An attempt to enable the capacity saving function failed. 
* `description`:(string) Short description about the volume. 
* `drive_type`:(string) Code indicating the drive type of the drive belonging to the volume. 
* `emulation_type`:(string) The volume emulation type or the volume status information.* `N/A` - Not available.* `NOT DEFINED` - The volume is not implemented.* `DEFINING` - The volume is being created.* `REMOVING` - The volume is being removed.* `OPEN-V` - To be provided by Hitachi. 
* `is_full_allocation_enabled`:(bool) Whether pages are reserved by the FullAllocation functionality. 
* `label`:(string) Label of the volume, as configured in the storage array. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `num_of_paths`:(int) Number of paths set for the volume. 
* `pool_id`:(string) ID of the pool with which the volume is associated. 
* `raid_level`:(string) RAID level for the volume.* `N/A` - RAID level is unknown or multiple RAID levels are being used.* `RAID1` - RAID1.* `RAID5` - RAID5.* `RAID6` - RAID6. 
* `raid_type`:(string) RAID type drive configuration. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `status`:(string) Status information of the volume.* `N/A` - The volume status is not available.* `NML` - The volume is in normal status.* `BLK` - The volume is in blocked state.* `BSY` - The volume status is being changed.* `Unknown` - The volume status is unknown (not supported). 
 