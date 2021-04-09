---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_aggregate"
description: |-
  NetApp aggregate is a collection of disks arranged into one or more RAID groups.
---

# Data Source: intersight_storage_net_app_aggregate
NetApp aggregate is a collection of disks arranged into one or more RAID groups.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_aggregate.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_type`:(string) Storage disk type for NetApp aggregate.* `HDD` - Hard Disk Drive.* `Hybrid` - Solid State Hard Disk Drive.* `Hybrid (Flash Pool)` - SSHD in a flash pool.* `SSD` - Solid State Disk.* `SSD (FabricPool)` - SSD in a flash pool.* `VMDisk (SDS)` - Storage disk with Hard disk drive.* `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives.* `LUN (FlexArray)` - LUN as a disk.* `Not Mapped` - Storage disk is not mapped. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Human readable name of the pool, limited to 64 characters. 
* `pool_id`:(string) Object ID for the pool. Platforms that use a number should convert it to string. 
* `raid_size`:(int) Size of the RAID group represented by number of disks in the group. 
* `raid_type`:(string) The RAID configuration type.* `Unknown` - Default unknown RAID type.* `RAID0` - RAID0 splits (\ stripes\ ) data evenly across two or more disks, without parity information.* `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.* `RAID4` - RAID4 stripes block level data and dedicates a disk to parity.* `RAID5` - RAID5  distributes striping and parity at a block level.* `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping.* `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.* `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.* `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the NetApp aggregate.* `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server.* `Online` - Aggregate is ready and available.* `Onlining` - Transitioning online.* `Offline` - Aggregate is unavailable.* `Offlining` - Transitioning offline.* `Relocating` - The aggregate is being relocated.* `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed.* `Failed` - The aggregate cannot be brought online.* `Inconsistent` - The aggregate has been marked corrupted; contact technical support.* `Unmounted` - The aggregate is not mounted. 
* `status`:(string) Human readable status of the pool, indicating the current health.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `type`:(string) Human readable type of the pool, such as thin, tiered, active-flash, etc. 
* `uuid`:(string) Uuid of  NetApp Aggregate. 
 