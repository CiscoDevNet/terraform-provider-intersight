---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup"
description: |-
        The Backup object is essential for managing system backups, ensuring data integrity and continuity through structured backup requests and processes.
        #### Purpose
        The Backup object oversees the creation and tracking of system backups, facilitating comprehensive data protection strategies through monitored backup operations.
        #### Key Concepts
        - **Managed States:** Enables tracking of backup states from initiation to completion, ensuring the integrity and availability of system data.
        - **Single Active Instance:** Maintains one active backup instance at any given time to streamline operations and avoid conflicts.
        - **Manual and Scheduled Modes:** Supports both manual and scheduled backup requests, accommodating diverse operational needs.
        - **Account Relationships:** Integrates with account management for secure and organized backup processes.

---

# Data Source: intersight_appliance_backup
The Backup object is essential for managing system backups, ensuring data integrity and continuity through structured backup requests and processes.
#### Purpose
The Backup object oversees the creation and tracking of system backups, facilitating comprehensive data protection strategies through monitored backup operations.
#### Key Concepts
- **Managed States:** Enables tracking of backup states from initiation to completion, ensuring the integrity and availability of system data.
- **Single Active Instance:** Maintains one active backup instance at any given time to streamline operations and avoid conflicts.
- **Manual and Scheduled Modes:** Supports both manual and scheduled backup requests, accommodating diverse operational needs.
- **Account Relationships:** Integrates with account management for secure and organized backup processes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_download_url`:(string) Download URL for the backup artifact when available. Only populated for successful local-protocol backups; empty for remote-protocol backups. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds since the backup process has started. 
* `end_time`:(string) End date and time of the backup process. 
* `filename`:(string) Backup filename to backup or restore. 
* `force_delete`:(bool) Set to true to allow deletion of the oldest local backup when local backup retention limit is reached. If false and retention count is reached, the backup operation fails. 
* `is_manual`:(bool) If true, represents a manual backup. Else represents a scheduled backup. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password to authenticate the fileserver. 
* `protocol`:(string) Communication protocol used by backup and restore workflow (e.g. scp, sftp, cifs, or local).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server.* `local` - Backup file is stored in Intersight Appliance. 
* `remote_host`:(string) Hostname of the remote file server. Not required when protocol is local. 
* `remote_path`:(string) File server directory or share name to copy the file. Not required when protocol is local. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). Not required when protocol is local. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date and time of the backup process. 
* `status`:(string) Status of the backup managed object.* `Started` - Backup or restore process has started.* `Created` - Backup or restore is in created state.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `Copied` - Backup file has been copied.* `Cleanup Failed` - Cleanup of the old backup has failed. 
* `use_policy_settings`:(bool) Set to true to inherit credentials, protocol, and file server settings from the appliance backup policy. If false, use explicit settings provided in this backup object. 
* `username`:(string) Username to authenticate the fileserver. Not required when protocol is local. 
 
