---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_rotate_data"
description: |-
        BackupRotateData tracks the set of current backups being managed under a backup rotation policy. It serves as the system’s record for actively managing backups during rotation, including when each backup was taken and which account scope it belongs to.
        #### Purpose
        Maintain rotation-aware backup bookkeeping so the platform can manage, retain, and rotate backups according to policy while providing read-only visibility into the backups currently under rotation management.
        #### Key Concepts
        - **Rotation management record**: Represents backups that are actively managed as part of a backup rotation process.
        - **Backup timestamping**: `backupTime` captures when the backup was taken, enabling ordering and retention/rotation decisions.
        - **Account-scoped ownership**: The `account` relationship ties the record to an `iam.Account`, and deletion of the account cascades to these records.
        - **Permission inheritance**: Inherits permissions from `account`, ensuring access aligns with account-level authorization.
        - **Read-only observability**: Exposed via READ for visibility; typically maintained by internal backup/rotation workflows rather than end-user CRUD operations.

---

# Data Source: intersight_appliance_backup_rotate_data
BackupRotateData tracks the set of current backups being managed under a backup rotation policy. It serves as the system’s record for actively managing backups during rotation, including when each backup was taken and which account scope it belongs to.
#### Purpose
Maintain rotation-aware backup bookkeeping so the platform can manage, retain, and rotate backups according to policy while providing read-only visibility into the backups currently under rotation management.
#### Key Concepts
- **Rotation management record**: Represents backups that are actively managed as part of a backup rotation process.
- **Backup timestamping**: `backupTime` captures when the backup was taken, enabling ordering and retention/rotation decisions.
- **Account-scoped ownership**: The `account` relationship ties the record to an `iam.Account`, and deletion of the account cascades to these records.
- **Permission inheritance**: Inherits permissions from `account`, ensuring access aligns with account-level authorization.
- **Read-only observability**: Exposed via READ for visibility; typically maintained by internal backup/rotation workflows rather than end-user CRUD operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup_rotate_data.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_time`:(string) The time at which the backup was taken. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) Backup filename to backup or restore. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protocol`:(string) Communication protocol used by backup and restore workflow (e.g. scp, sftp, cifs, or local).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server.* `local` - Backup file is stored in Intersight Appliance. 
* `remote_host`:(string) Hostname of the remote file server. Not required when protocol is local. 
* `remote_path`:(string) File server directory or share name to copy the file. Not required when protocol is local. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). Not required when protocol is local. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) Username to authenticate the fileserver. Not required when protocol is local. 
 
