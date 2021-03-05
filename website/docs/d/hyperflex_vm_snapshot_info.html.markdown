---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vm_snapshot_info"
description: |-
  Virtual Machine Snapshot information like replication status, snapshot point and status.
---

# Data Source: intersight_hyperflex_vm_snapshot_info
Virtual Machine Snapshot information like replication status, snapshot point and status.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_snapshot_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `label`:(string) The name of the Virtual Machine and the time stamp of the snapshot. 
* `mode`:(string) Quiesce Mode for snapshot taken on Hyperflex cluster.* `NONE` - The snapshot quiesce mode is none.* `CRASH` - The snapshot quiesce mode is crash.* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `replication_status`:(string) Replication status of the least successful target cluster.* `NONE` - Snapshot replication state is none.* `SUCCESS` - Snapshot completed successfully.* `FAILED` - Snapshot failed replication status code.* `PAUSED` - Snapshot replication paused status code.* `IN_USE` - Snapshot replica in use status code.* `STARTING` - Snapshot replication starting.* `REPLICATING` - Snapshot replication in progress. 
* `snapshot_status`:(string) Snapshot status of the source cluster.* `SUCCESS` - This snapshot status code is success.* `FAILED` - This snapshot status code is failed.* `IN_PROGRESS` - This snapshot status code is in progress.* `DELETING` - This snapshot status code is deleting.* `DELETED` - This snapshot status code is deleted.* `NONE` - This snapshot status code is none. 
* `source_timestamp`:(int) Timestamp the snapshot was created on the source cluster. 
 