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
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_snapshot_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `display_status`:(string) Combined status code from replication and snapshot status to display in the Intersight UI.* `NONE` - Snapshot replication state is none.* `SUCCESS` - Snapshot completed successfully.* `FAILED` - Snapshot failed replication status code.* `PAUSED` - Snapshot replication paused status code.* `IN_USE` - Snapshot replica in use status code.* `STARTING` - Snapshot replication starting.* `REPLICATING` - Snapshot replication in progress. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) The name of the Virtual Machine and the time stamp of the snapshot. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Quiesce Mode for snapshot taken on Hyperflex cluster.* `NONE` - The snapshot quiesce mode is none.* `CRASH` - The snapshot quiesce mode is crash.* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `replication_status`:(string) Replication status of the least successful target cluster.* `NONE` - Snapshot replication state is none.* `SUCCESS` - Snapshot completed successfully.* `FAILED` - Snapshot failed replication status code.* `PAUSED` - Snapshot replication paused status code.* `IN_USE` - Snapshot replica in use status code.* `STARTING` - Snapshot replication starting.* `REPLICATING` - Snapshot replication in progress. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_error_msg`:(string) Error message from snapshot creation or replcation if any exist. 
* `snapshot_status`:(string) Snapshot status of the source cluster.* `SUCCESS` - This snapshot status code is success.* `FAILED` - This snapshot status code is failed.* `IN_PROGRESS` - This snapshot status code is in progress.* `DELETING` - This snapshot status code is deleting.* `DELETED` - This snapshot status code is deleted.* `NONE` - This snapshot status code is none.* `INIT` - This snapshot status code is initializing. 
* `source_timestamp`:(int) Timestamp the snapshot was created on the source cluster. 
* `src_cluster_name`:(string) Name of the cluster this snapshot resides on. 
* `target_completion_timestamp`:(int) Timestamp the snapshot was finished replicating on the target cluster. 
* `tgt_cluster_name`:(string) Name of the cluster this snapshot is replicated to. 
 
