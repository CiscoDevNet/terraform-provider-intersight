---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_monitor"
description: |-
        The BackupMonitor object plays a pivotal role in overseeing the history and status of backup operations, ensuring ongoing data protection and system reliability.
        Purpose
        BackupMonitor tracks backup activities, providing a clear overview of the backup status and alerting users to potential issues or outdated backups.
        #### Key Concepts
        - **Status Tracking:** Monitors the success, failure, and currency of backups, offering timely feedback on data protection health.
        - **Alarm Integration:** Works in conjunction with alarm definitions to alert users to backup failures or outdated processes, ensuring proactive management.
        - **System-Oriented:** Designed for automated, system-driven monitoring, enhancing operational efficiency and reliability.
        - **Account Integration:** Connects seamlessly with account structures for consistent and secure monitoring operations.

---

# Data Source: intersight_appliance_backup_monitor
The BackupMonitor object plays a pivotal role in overseeing the history and status of backup operations, ensuring ongoing data protection and system reliability.
Purpose
BackupMonitor tracks backup activities, providing a clear overview of the backup status and alerting users to potential issues or outdated backups.
#### Key Concepts
- **Status Tracking:** Monitors the success, failure, and currency of backups, offering timely feedback on data protection health.
- **Alarm Integration:** Works in conjunction with alarm definitions to alert users to backup failures or outdated processes, ensuring proactive management.
- **System-Oriented:** Designed for automated, system-driven monitoring, enhancing operational efficiency and reliability.
- **Account Integration:** Connects seamlessly with account structures for consistent and secure monitoring operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup_monitor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) Filename of the backup for the backup monitor. 
* `last_backup_rotation_status`:(string) Status of the oldest Intersight Appliance backup cleanup.* `BackupFound` - Backup is successful and complete.* `BackupFailed` - The current Backup failed.* `BackupOutdated` - Backup is old and outdated.* `BackupCleanupFailed` - Cleanup of the old backup has failed. 
* `last_backup_status`:(string) Status of the most recent Intersight Appliance backup.* `BackupFound` - Backup is successful and complete.* `BackupFailed` - The current Backup failed.* `BackupOutdated` - Backup is old and outdated.* `BackupCleanupFailed` - Cleanup of the old backup has failed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
