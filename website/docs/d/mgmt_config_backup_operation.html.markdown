---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_backup_operation"
description: |-
        The ConfigBackupOperation object represents an administrator-initiated configuration backup run. This provides a single, traceable record for when a backup was requested, how it was executed, and how its progress and completion are tracked end-to-end.
        #### Purpose
        ConfigBackupOperation exists to orchestrate and audit the lifecycle of creating a configuration backup for an account. It acts as the top-level container that coordinates organization-level and account-level export activities and ties them to a resulting backup artifact.
        #### Key Concepts
        - **Lifecycle Tracking:** Captures the operational state of a backup from initiation through completion, enabling status-based monitoring and user visibility.
        - **Account-Scoped Execution:** Runs within a specific account context, ensuring backups align with tenant boundaries and permissions.
        - **Coordinated Export Workflow:** Aggregates export activity metadata so customers can understand what was included and which export jobs were triggered.
        - **Retention and Rollover Awareness:** Supports controlled space management via rollover behavior and retention protection, enabling predictable backup management at scale.
        - **Artifact Association:** Links the operation to the backup instance produced (the durable item used for download and restore workflows).

---

# Data Source: intersight_mgmt_config_backup_operation
The ConfigBackupOperation object represents an administrator-initiated configuration backup run. This provides a single, traceable record for when a backup was requested, how it was executed, and how its progress and completion are tracked end-to-end.
#### Purpose
ConfigBackupOperation exists to orchestrate and audit the lifecycle of creating a configuration backup for an account. It acts as the top-level container that coordinates organization-level and account-level export activities and ties them to a resulting backup artifact.
#### Key Concepts
- **Lifecycle Tracking:** Captures the operational state of a backup from initiation through completion, enabling status-based monitoring and user visibility.
- **Account-Scoped Execution:** Runs within a specific account context, ensuring backups align with tenant boundaries and permissions.
- **Coordinated Export Workflow:** Aggregates export activity metadata so customers can understand what was included and which export jobs were triggered.
- **Retention and Rollover Awareness:** Supports controlled space management via rollover behavior and retention protection, enabling predictable backup management at scale.
- **Artifact Association:** Links the operation to the backup instance produced (the durable item used for download and restore workflows).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_backup_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_account`:(string) The moid of the account from which the backup was taken. 
* `backup_account_domain_group_moid`:(string) The domain group moid of the account from which the backup was taken. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the backup operation. 
* `destination`:(string) Indicates whether the backup was created locally or imported.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) End date and time of the backup operation. 
* `is_aes_key_set`:(bool) Indicates whether the value of the 'aesKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided identifier for the backup operation. 
* `retention_lock`:(bool) When set, ensures that the backup archive is protected from deletion and rollover operations. The value for retention lock is in sync with the backup instance created as part of this operation. 
* `rollover_initiated`:(bool) Internal property used to indicate that a rollover has been initiated as part of this operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date and time of the backup operation. 
* `status`:(string) The current status of the backup operation.* `Scheduled` - Backup or restore process has been scheduled.* `NotInitiated` - Backup or restore process has not been initiated.* `InProgress` - Backup or restore process is in progress.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `NotComplete` - Backup or restore process has partially imported configuration. 
* `system_admin_triggered`:(bool) Set when a system administrator initiated the backup operation. 
* `user_password`:(string) The password provided by the user to decrypt a password encrypted backup file during import operation. This password is used to decrypt the backup archive before validation when importing a password encrypted backup.This field is optional and only required when importing a password encrypted backup file. 
 
