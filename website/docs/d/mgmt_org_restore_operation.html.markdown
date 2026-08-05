---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_org_restore_operation"
description: |-
        An organization-scoped portion of a configuration restore run with a dedicated view into how a restore is applied per organization, including sequencing and status.
        #### Purpose
        OrgRestoreOperation breaks down a restore operation into organization-level units so that organization-owned configuration can be processed and tracked independently while still being governed by the parent restore operation.
        #### Key Concepts
        - **Per-Organization Restore Tracking:** Enables fine-grained monitoring of restore progress and outcomes per organization.
        - **Sequenced Processing:** Supports deterministic ordering of organization restores to ensure dependencies and shared contexts organizations are handled predictably.
        - **Inherited Governance and Context:** Operates under the authorization and operational intent of the parent restore operation for consistency and auditability.
        - **Import Job Alignment:** Connects organization restore activity to the underlying import mechanisms used to apply configuration changes.

---

# Data Source: intersight_mgmt_org_restore_operation
An organization-scoped portion of a configuration restore run with a dedicated view into how a restore is applied per organization, including sequencing and status.
#### Purpose
OrgRestoreOperation breaks down a restore operation into organization-level units so that organization-owned configuration can be processed and tracked independently while still being governed by the parent restore operation.
#### Key Concepts
- **Per-Organization Restore Tracking:** Enables fine-grained monitoring of restore progress and outcomes per organization.
- **Sequenced Processing:** Supports deterministic ordering of organization restores to ensure dependencies and shared contexts organizations are handled predictably.
- **Inherited Governance and Context:** Operates under the authorization and operational intent of the parent restore operation for consistency and auditability.
- **Import Job Alignment:** Connects organization restore activity to the underlying import mechanisms used to apply configuration changes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_org_restore_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `org_name`:(string) Name of the organization for which the restore operation is being performed. 
* `sequence_number`:(int) The sequence number that determines the order in which organization configurations will be processed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The current status of the restore operation.* `Scheduled` - Backup or restore process has been scheduled.* `NotInitiated` - Backup or restore process has not been initiated.* `InProgress` - Backup or restore process is in progress.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `NotComplete` - Backup or restore process has partially imported configuration. 
 
