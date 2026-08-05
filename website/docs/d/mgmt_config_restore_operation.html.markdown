---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_restore_operation"
description: |-
        The ConfigRestoreOperation object represents an administrator-initiated configuration restore run. It tracks the end-to-end process of applying a previously captured backup to a target account, including progress, outcomes, and workflow coordination.
        #### Purpose
        ConfigRestoreOperation exists to orchestrate restoration of configuration from a chosen backup source, providing a top-level record that drives restore execution, tracks status, and reports import progress across account and organization scopes.
        #### Key Concepts
        - **End-to-End Restore Orchestration:** Coordinates staging, importing, and completion of a restore run, enabling customers to manage restoration as a single operation.
        - **Source Flexibility:** Designed to support restoring from different backup locations (e.g., local repository), depending on operational needs.
        - **Controlled Restore Behavior:** Supports configurable restore strategies and error-handling behavior so customers can align restores with change-management requirements.
        - **Progress and Outcome Visibility:** Aggregates import status reporting so customers can monitor completion and detect partial or incomplete outcomes.
        - **Workflow Integration:** Links to underlying workflow execution details to support operational auditability and troubleshooting.

---

# Data Source: intersight_mgmt_config_restore_operation
The ConfigRestoreOperation object represents an administrator-initiated configuration restore run. It tracks the end-to-end process of applying a previously captured backup to a target account, including progress, outcomes, and workflow coordination.
#### Purpose
ConfigRestoreOperation exists to orchestrate restoration of configuration from a chosen backup source, providing a top-level record that drives restore execution, tracks status, and reports import progress across account and organization scopes.
#### Key Concepts
- **End-to-End Restore Orchestration:** Coordinates staging, importing, and completion of a restore run, enabling customers to manage restoration as a single operation.
- **Source Flexibility:** Designed to support restoring from different backup locations (e.g., local repository), depending on operational needs.
- **Controlled Restore Behavior:** Supports configurable restore strategies and error-handling behavior so customers can align restores with change-management requirements.
- **Progress and Outcome Visibility:** Aggregates import status reporting so customers can monitor completion and detect partial or incomplete outcomes.
- **Workflow Integration:** Links to underlying workflow execution details to support operational auditability and troubleshooting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_config_restore_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_account_moid`:(string) The moid of the account from which the backup was taken. 
* `backup_name`:(string) Name of the backup that needs to be restored. 
* `backup_source`:(string) Indicates whether the backup was created locally or imported.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) End date and time of the restore process. 
* `ignore_secure_properties`:(bool) If set to true, secure properties will not be decrypted, and empty values will be used for these properties. As a result, if a secure property field is mandatory, the MO will be marked as incomplete. 
* `is_aes_key_set`:(bool) Indicates whether the value of the 'aesKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided identifier for the restore operation. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_source`:(string) The type of source from which the backup - local or remote - needs to be restored.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `start_time`:(string) Start date and time of the restore process. 
* `status`:(string) Status of the restore operation.* `Scheduled` - Backup or restore process has been scheduled.* `NotInitiated` - Backup or restore process has not been initiated.* `InProgress` - Backup or restore process is in progress.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `NotComplete` - Backup or restore process has partially imported configuration. 
 
