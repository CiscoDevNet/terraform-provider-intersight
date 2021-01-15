---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_backup_policy"
description: |-
  Specifies cluster backup configuration for a HyperFlex Cluster.
---

# Data Source: intersight_hyperflex_cluster_backup_policy
Specifies cluster backup configuration for a HyperFlex Cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_data_store_name`:(string) Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up. 
* `backup_data_store_size`:(int) Replication data store size in backupDataStoreSizeUnit. 
* `backup_data_store_size_unit`:(string) Replication data store size. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `replication_pair_name_prefix`:(string) Replication cluster pairing name prefix. 
* `snapshot_retention_count`:(int) Number of snapshots that will be retained as part of the Multi Point in Time support. 
