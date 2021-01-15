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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_status`:(string) Description of the backup status of this VmBackupInfo.* `InitializingProtection` - Protection has started, but not completed.* `Protected` - Protection has completed successfully.* `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protection_status`:(string) Description of the protection status of this VmBackupInfo.* `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started.* `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed.* `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed.* `FAILOVER_STARTED` - The protection status is failover started.* `FAILOVER_FAILED` - The protection status is failover failed.* `FAILOVER_COMPLETED` - The protection status is failover completed.* `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started.* `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed.* `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed.* `REVERSEPROTECT_STARTED` - The protection status is reverse protect started.* `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed.* `ACTIVE` - The protection status is active.* `CREATION_IN_PROGRESS` - The protection status is failover in progress.* `CREATION_FAILED` - The protection status is creation failed.* `LOCAL_RESTORE_STARTED` - The protection status is local restore started.* `LOCAL_RESTORE_FAILED` - The protection status is local restore failed. 
