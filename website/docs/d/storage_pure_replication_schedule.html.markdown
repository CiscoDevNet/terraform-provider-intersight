---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_replication_schedule"
description: |-
  Pure snapshot replication schedule entity.
---

# Data Source: intersight_storage_pure_replication_schedule
Pure snapshot replication schedule entity.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `daily_limit`:(int) Total number of snapshots per day to be available on target above and over the specified retention time. PureStorageFlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired.In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day. 
* `frequency`:(string) Replication frequency. It is an interval at which replication is set to trigger.Examples:    PT2H, Snapshot is generated every 2 hours.    P30D, Snapshot is scheduled for every 30 days.    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Replication schedule name. 
* `replication_time`:(string) Preferred time of day at which to replicate the snapshots on target array.It is applicable only if the replication frequency is set for a day or more.Format: hh:mm:ssExample: 15:00:00, Replication is set for 3:00 PM. 
* `retention_time`:(string) Duration to keep the replicated snapshots on the targets.Replicated snapshots are deleted from target array once the retention period has elapsed.Examples:P30D, Snapshots are available for 30 days.PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds. 
* `snapshot_expiry_time`:(string) Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period. 
