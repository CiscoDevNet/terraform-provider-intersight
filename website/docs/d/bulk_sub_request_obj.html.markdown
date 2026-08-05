---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_sub_request_obj"
description: |-
        The SubRequestObj object represents an individual API request processed as part of a larger bulk operation. It serves as an internal record for every sub-request included in an incoming bulk request.
        #### Purpose
        The SubRequestObj is designed to provide granular tracking and status reporting for each specific operation (Create, Update, or Delete) within a bulk request. It allows the system to manage, monitor, and correlate the execution status of multiple sub-requests, ensuring that each part of a bulk operation is accounted for and processed correctly.
        #### Key Concepts
        - **Request Correlation:** Uses a unique request number to correlate individual sub-requests with their corresponding bulk request and response.
        - **Execution Tracking:** Monitors the lifecycle of each sub-request, including execution start/completion times and current processing status.
        - **Presence Detection:** Includes flags to detect if an object is already present or if a system-defined object has been triggered, supporting idempotent operations.
        - **Result Management:** Stores the outcome of the individual API action, including success or failure details, within the `result` property.
        - **Target Context:** Maintains the target MOID and request body, ensuring that the operation is applied to the correct resource with the intended data payload.

---

# Data Source: intersight_bulk_sub_request_obj
The SubRequestObj object represents an individual API request processed as part of a larger bulk operation. It serves as an internal record for every sub-request included in an incoming bulk request.
#### Purpose
The SubRequestObj is designed to provide granular tracking and status reporting for each specific operation (Create, Update, or Delete) within a bulk request. It allows the system to manage, monitor, and correlate the execution status of multiple sub-requests, ensuring that each part of a bulk operation is accounted for and processed correctly.
#### Key Concepts
- **Request Correlation:** Uses a unique request number to correlate individual sub-requests with their corresponding bulk request and response.
- **Execution Tracking:** Monitors the lifecycle of each sub-request, including execution start/completion times and current processing status.
- **Presence Detection:** Includes flags to detect if an object is already present or if a system-defined object has been triggered, supporting idempotent operations.
- **Result Management:** Stores the outcome of the individual API action, including success or failure details, within the `result` property.
- **Target Context:** Maintains the target MOID and request body, ensuring that the operation is applied to the correct resource with the intended data payload.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_sub_request_obj.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `body_string`:(string) The body of the sub-request in string format. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execution_completion_time`:(string) The time at which processing of this request completed. 
* `execution_start_time`:(string) The time at which processing of this request started. 
* `is_bulk_mo_op`:(bool) For Async Bulk Mo Operations this flag will be set to true. 
* `is_object_present`:(bool) This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_duplicates`:(bool) Skip the already present objects. The value from the Request. 
* `status`:(string) The status of the request.* `Pending` - Indicates that the request is yet to be processed.* `ObjPresenceCheckInProgress` - Indicates that the checking for object presence is in progress.* `ObjPresenceCheckInComplete` - Indicates that the request is being processed.* `ObjPresenceCheckFailed` - Indicates that the checking for object presence failed.* `Processing` - Indicates that the request is being processed.* `TimedOut` - Indicates that the request processing timed out.* `Failed` - Indicates that the request processing failed.* `Completed` - Indicates that the request processing is complete.* `Partial` - Indicates that the request is partially complete due to an error.* `Skipped` - Indicates that the request was skipped. 
* `system_defined_object_detected`:(bool) This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence. 
* `target_moid`:(string) Used with PATCH & DELETE actions. The moid of an existing object instance. 
* `uri`:(string) The URI on which this bulk action is to be performed. 
* `verb`:(string) The type of operation to be performed.One of - Post (Create), Patch (Update) or Delete (Remove).* `POST` - Used to create a REST resource.* `PATCH` - Used to update a REST resource.* `DELETE` - Used to delete a REST resource. 
 
