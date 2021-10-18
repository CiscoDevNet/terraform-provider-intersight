---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_export"
description: |-
  All export operations are captured as Export instances. Users shall use this Export
mo to track the export operation progress.
---

# Data Source: intersight_bulk_export
All export operations are captured as Export instances. Users shall use this Export
mo to track the export operation progress.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_export.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be performed on the export operation.* `Start` - Starts the export operation.* `Cancel` - Cancels the export operation that is in progress. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `export_tags`:(bool) Specifies whether tags must be exported and will be considered for all the items MOs. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the export operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `OrderInProgress` - The archive operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `OperationTimedOut` - The operation has timed out.* `OperationCancelled` - The operation has been cancelled.* `CancelInProgress` - The operation is being cancelled. 
* `status_message`:(string) Status message associated with failures or progress indication. 
 