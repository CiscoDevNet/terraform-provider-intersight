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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_backup_policy_deployment.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_data_store_name`:(string) Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up. 
* `backup_data_store_size`:(int) Replication data store size in backupDataStoreSizeUnit. 
* `backup_data_store_size_unit`:(string) Replication data store size. 
* `create_time`:(string) The time when this managed object was created. 
* `data_store_encryption_enabled`:(bool) Whether the datastore is encrypted or not. 
* `description`:(string) Description from corresponding ClusterBackupPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `local_snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name from corresponding ClusterBackupPolicy. 
* `policy_moid`:(string) Deployed cluster policy moid. 
* `profile_moid`:(string) Deployed cluster profile moid. 
* `replication_pair_name_prefix`:(string) Replication cluster pairing name prefix. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
* `source_detached`:(bool) True if policy was detached from source Hyperflex Cluster. 
* `source_request_id`:(string) Unique source cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `source_uuid`:(string) Uuid of the source Hyperflex Cluster. 
* `target_detached`:(bool) True if policy was detached from target Hyperflex Cluster. 
* `target_request_id`:(string) Unique target cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `target_uuid`:(string) Uuid of the target Hyperflex Cluster. 
 