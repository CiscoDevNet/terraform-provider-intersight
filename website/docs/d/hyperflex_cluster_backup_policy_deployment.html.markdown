---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_backup_policy_deployment"
description: |-
  Record of HyperFlex Cluster backup policy deployment.
---

# Data Source: intersight_hyperflex_cluster_backup_policy_deployment
Record of HyperFlex Cluster backup policy deployment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_data_store_name`:(string) Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up. 
* `backup_data_store_size`:(int) Replication data store size in backupDataStoreSizeUnit. 
* `backup_data_store_size_unit`:(string) Replication data store size. 
* `description`:(string) Description from corresponding ClusterBackupPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name from corresponding ClusterBackupPolicy. 
* `policy_moid`:(string) Deployed cluster policy moid. 
* `profile_moid`:(string) Deployed cluster profile moid. 
* `replication_pair_name_prefix`:(string) Replication cluster pairing name prefix. 
* `snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
* `source_detached`:(bool) True if policy was detached from source Hyperflex Cluster. 
* `source_request_id`:(string) Unique source cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `source_uuid`:(string) Uuid of the source Hyperflex Cluster. 
* `target_detached`:(bool) True if policy was detached from target Hyperflex Cluster. 
* `target_request_id`:(string) Unique target cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `target_uuid`:(string) Uuid of the target Hyperflex Cluster. 
