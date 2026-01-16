---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_exported_item"
description: |-
        A single managed object that is being exported.

---

# Data Source: intersight_bulk_exported_item
A single managed object that is being exported.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_exported_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `exclude_relations`:(bool) Do not export relationships. 
* `export_tags`:(bool) Specifies whether tags must be exported for item MO. 
* `file_name`:(string) Name of the file corresponding to item MO. 
* `include_org_identity`:(bool) Indicates that exported references for objects which are organization owned should include the organization reference along with the other identity properties. 
* `is_aes_key_set`:(bool) Indicates whether the value of the 'aesKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) MO item identity (the moref corresponding to item) expressed as a string. 
* `service_name`:(string) Name of the target service that owns the item MO. Service responsible for handling exported item mo notifications. 
* `service_version`:(string) Version of the service that owns the item MO. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the item's export operation.* `` - The operation has not started.* `Ready` - The operation is ready to start.* `ValidationInProgress` - The validation operation is in progress.* `Valid` - The content to be imported is valid.* `InValid` - The content to be imported is not valid and the status message will have the reason.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been initiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out.* `OperationCancelled` - The operation has been canceled.* `CancelInProgress` - The operation is being canceled. 
* `status_message`:(string) Progress or error message for the MO's export operation. 
 
