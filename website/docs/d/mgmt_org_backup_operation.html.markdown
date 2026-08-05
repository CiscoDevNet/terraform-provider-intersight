---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_org_backup_operation"
description: |-
        The OrgBackupOperation object represents the organization-scoped portion of a broader configuration backup run. This provides visibility into how backup processing is applied per organization, while remaining governed by the parent backup operation.
        #### Purpose
        OrgBackupOperation exists to segment backup execution by organization so that organization-owned configuration can be handled, tracked, and reported exported independently within a single account-level backup.
        #### Key Concepts
        - **Per-Organization Backup Tracking:** Enables monitoring and reporting at the organization boundary, useful for environments with multiple orgs and delegated management.
        - **Inherited Governance:** Leverages the access controls and execution context of the parent backup operation to maintain consistent authorization and auditability.
        - **Category-Centric Reporting:** Organizes export activity by configuration categories to improve transparency of what was processed under an organization.
        - **Export Job Alignment:** Connects organization backups to the underlying export mechanisms used to produce the backup content.
        A subset of the config backup operation triggered by an admin user which creates configuration backups for an organization.

---

# Data Source: intersight_mgmt_org_backup_operation
The OrgBackupOperation object represents the organization-scoped portion of a broader configuration backup run. This provides visibility into how backup processing is applied per organization, while remaining governed by the parent backup operation.
#### Purpose
OrgBackupOperation exists to segment backup execution by organization so that organization-owned configuration can be handled, tracked, and reported exported independently within a single account-level backup.
#### Key Concepts
- **Per-Organization Backup Tracking:** Enables monitoring and reporting at the organization boundary, useful for environments with multiple orgs and delegated management.
- **Inherited Governance:** Leverages the access controls and execution context of the parent backup operation to maintain consistent authorization and auditability.
- **Category-Centric Reporting:** Organizes export activity by configuration categories to improve transparency of what was processed under an organization.
- **Export Job Alignment:** Connects organization backups to the underlying export mechanisms used to produce the backup content.
A subset of the config backup operation triggered by an admin user which creates configuration backups for an organization.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_mgmt_org_backup_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `org_name`:(string) Name of the organization to be backed up. 
* `preserve_identities`:(bool) The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
