---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_export"
description: |-
        The Export object is integral to managing export operations within the system, providing a structured approach to monitor and control various aspects of data exporting activities.
        #### Purpose
        The Export object tracks the progress and status of data export operations. It serves as the main point of reference for users who need to initiate, manage, or terminate export processes efficiently.
        #### Key Concepts
        - **Operation Tracking:** Captures all export activities as instances, allowing users to monitor status updates and completion metrics.
        - **Privilege Management:** Enforces access control through privilege sets, ensuring only authorized users can initiate, update, or delete export processes.
        - **Action Flexibility:** Supports actions such as starting or canceling an export, giving users control over operation flow.
        - **Organizational Ownership:** Reflects ownership and management by the organization, promoting structured and accountable export operations.

---

# Data Source: intersight_bulk_export
The Export object is integral to managing export operations within the system, providing a structured approach to monitor and control various aspects of data exporting activities.
#### Purpose
The Export object tracks the progress and status of data export operations. It serves as the main point of reference for users who need to initiate, manage, or terminate export processes efficiently.
#### Key Concepts
- **Operation Tracking:** Captures all export activities as instances, allowing users to monitor status updates and completion metrics.
- **Privilege Management:** Enforces access control through privilege sets, ensuring only authorized users can initiate, update, or delete export processes.
- **Action Flexibility:** Supports actions such as starting or canceling an export, giving users control over operation flow.
- **Organizational Ownership:** Reflects ownership and management by the organization, promoting structured and accountable export operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_export.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be performed on the export operation.* `Start` - Starts the export operation.* `Cancel` - Cancels the export operation that is in progress. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `exclude_relations`:(bool) Used to specify that none of the relationships should be exported. 
* `export_tags`:(bool) Specifies whether tags must be exported and will be considered for all the items MOs. 
* `include_org_identity`:(bool) Indicates that exported references for objects which are organization owned should include the organization reference along with the other identity properties. 
* `is_aes_key_set`:(bool) Indicates whether the value of the 'aesKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). 
* `permission_id`:(string) The permission identifier which indicates the permission that current user has that will allow to start this export operation. 
* `preserve_identities`:(bool) The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `split_requests`:(bool) Intersight supports the import operation using bulk requests. A single bulk request is created using the bulk subrequests created in export operation. There is a restriction of a single API request size in Intersight.If the exported MOs are many in an export operation, this flag will store the bulk sub requests as two dimensionalcollection property exportedObjectsAsGroups instead of one dimensional exportedObjects value where the sub requests are split into multiple smallergroups. Each group can be sent in a single bulk request during import operation. 
* `status`:(string) Status of the export operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `OrderInProgress` - The archive operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `OperationTimedOut` - The operation has timed out.* `OperationCancelled` - The operation has been cancelled.* `CancelInProgress` - The operation is being cancelled. 
* `status_message`:(string) Status message associated with failures or progress indication. 
* `user_id`:(string) The user identifier which indicates the user that started this export operation. 
 
