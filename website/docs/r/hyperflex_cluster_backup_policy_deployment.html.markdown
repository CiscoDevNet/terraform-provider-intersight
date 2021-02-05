---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_backup_policy_deployment"
description: |-
  Record of HyperFlex Cluster backup policy deployment.
---

# Resource: intersight_hyperflex_cluster_backup_policy_deployment
Record of HyperFlex Cluster backup policy deployment.
## Argument Reference
The following arguments are supported:
* `backup_data_store_name`:(string)(Computed) Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up. 
* `backup_data_store_size`:(int)(Computed) Replication data store size in backupDataStoreSizeUnit. 
* `backup_data_store_size_unit`:(string)(Computed) Replication data store size. 
* `backup_target`:(HashMap) - A reference to a hyperflexCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string)(Computed) Description from corresponding ClusterBackupPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string)(Computed) Name from corresponding ClusterBackupPolicy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_moid`:(string)(Computed) Deployed cluster policy moid. 
* `profile_moid`:(string)(Computed) Deployed cluster profile moid. 
* `replication_pair_name_prefix`:(string)(Computed) Replication cluster pairing name prefix. 
* `replication_schedule`:(HashMap) -(Computed) Backup policy replication schedule. 
This complex property has following sub-properties:
  + `backup_interval`:(int) Time interval between two copies in minutes. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `snapshot_retention_count`:(int)(Computed) Number of snapshots that will be retained as part of the Multi Point in Time support. 
* `source_cluster`:(HashMap) - A reference to a hyperflexCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `source_detached`:(bool) True if policy was detached from source Hyperflex Cluster. 
* `source_request_id`:(string)(Computed) Unique source cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `source_uuid`:(string)(Computed) Uuid of the source Hyperflex Cluster. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_detached`:(bool) True if policy was detached from target Hyperflex Cluster. 
* `target_request_id`:(string)(Computed) Unique target cluster request ID allowing retry of the same logical request following a transient communication failure. 
* `target_uuid`:(string)(Computed) Uuid of the target Hyperflex Cluster. 


## Import
`intersight_hyperflex_cluster_backup_policy_deployment` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_cluster_backup_policy_deployment.example 1234567890987654321abcde
``` 