---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_host_utility_operation"
description: |-
        Host operation that need to be performed using host utility (HSU), like secure erase, secure erase with decommission and scrub are managed by this MO.

---

# Data Source: intersight_compute_host_utility_operation
Host operation that need to be performed using host utility (HSU), like secure erase, secure erase with decommission and scrub are managed by this MO.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_host_utility_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_utility_operation_mode`:(string) Host utility operation need to be performed in the endpoint.* `None` - Host utility mode of the operation is set to none by default.* `SecureErase` - EU LOT-9 secure data cleanup on the server components.* `SecureEraseWithDecommission` - EU LOT-9 secure data cleanup on the server components and do decommission.* `Scrub` - Quick cleanup on storage and BIOS. 
* `host_utility_operation_status`:(string) Task status of the host utility operation.* `Initiated` - This status indicates that host utility operation request is initiated.* `InProgress` - The operation status indicates that host utility operation is in-progress after the basic validations.* `CompletedOk` - The operation status indicates that host utility operation is completed successfully with no error or warning.* `CompletedError` - The operation status indicates that host utility operation is completed with error.* `CompletedWarning` - The operation status indicates that host utility operation is completed with warning.* `Aborted` - The operation status indicates that host utility operation is terminated or aborted.* `Invalidated` - The operation status indicates that host utility operation is invalid due to validation failure. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
