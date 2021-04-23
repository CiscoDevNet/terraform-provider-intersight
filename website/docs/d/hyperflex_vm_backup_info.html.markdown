---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vm_backup_info"
description: |-
  Virtual Machine backup information.
---

# Data Source: intersight_hyperflex_vm_backup_info
Virtual Machine backup information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_backup_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_status`:(string) Description of the backup status of this VmBackupInfo.* `InitializingProtection` - Protection has started, but not completed.* `Protected` - Protection has completed successfully.* `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protection_status`:(string) Description of the protection status of this VmBackupInfo.* `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started.* `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed.* `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed.* `FAILOVER_STARTED` - The protection status is failover started.* `FAILOVER_FAILED` - The protection status is failover failed.* `FAILOVER_COMPLETED` - The protection status is failover completed.* `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started.* `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed.* `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed.* `REVERSEPROTECT_STARTED` - The protection status is reverse protect started.* `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed.* `ACTIVE` - The protection status is active.* `CREATION_IN_PROGRESS` - The protection status is failover in progress.* `CREATION_FAILED` - The protection status is creation failed.* `LOCAL_RESTORE_STARTED` - The protection status is local restore started.* `LOCAL_RESTORE_FAILED` - The protection status is local restore failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 