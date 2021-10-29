---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_backup_policy_inventory"
description: |-
  Response to Backup Policy requests and queries.
---

# Data Source: intersight_hyperflex_cluster_backup_policy_inventory
Response to Backup Policy requests and queries.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_backup_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Validate, Deploy, Prepare, Commit, or Abort the Backup Policy. Allowed values are \ VALIDATE\ , \ DEPLOY\ , \ PREPARE\ , \ COMMIT\ , \ ABORT\ .* `VALIDATE` - Check for problems in policy request without deploying.* `DEPLOY` - Remove the policy.  Only allowed when cleanup field is true.* `PREPARE` - Prepare the policy for subsequent \ COMMIT\  or \ ABORT\ .  Only allowed when cleanup field is false.* `COMMIT` - Commit the prepared policy.  Only allowed when cleanup field is false.* `ABORT` - Abort the prepared policy.  Only allowed when cleanup field is false. 
* `cleanup`:(bool) If true, remove the policy. Otherwise proceed with the specified policy action. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_source`:(bool) Indicates if the HyperFlex Cluster is the source cluster or the target cluster. When set to true, the HyperFlex Cluster is the source for VM backups. When set to false, the HyperFlex Cluster is the target cluster where VM backups are saved. 
* `job_details`:(string) Details from associated HyperFlex job execution. 
* `job_exception_message`:(string) Job Exception message from associated HyperFlex job execution. 
* `job_id`:(string) Job ID of associated HyperFlex job. 
* `job_state`:(string) State of the associated HyperFlex job. When present possible values are \ RUNNING\ , \ COMPLETED\  or \ EXCEPTION\ .* `RUNNING` - HyperFlex policy job is currently \ RUNNING\ .* `COMPLETED` - HyperFlex policy job completed successfully.* `EXCEPTION` - HyperFlex policy job failed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `policy_moid`:(string) Intersight HyperFlex Cluster Backup Policy MOID. 
* `request_id`:(string) Unique request ID allowing retry of the same logical request following a transient communication failure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_uuid`:(string) UUID of the source HyperFlex Cluster. 
* `target_uuid`:(string) UUID of the target HyperFlex Cluster. 
* `nr_version`:(int) Version of the Backup Policy. 
 