---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_result"
description: |-
        The Result API is a status-monitor resource used to show the processing status of any bulk MO API when the request HTTP 'prefer' header is set to 'respond-async' value.

---

# Data Source: intersight_bulk_result
The Result API is a status-monitor resource used to show the processing status of any bulk MO API when the request HTTP 'prefer' header is set to 'respond-async' value.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_result.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action_on_error`:(string) The action that will be performed when an error occurs during processing of the request.* `Stop` - Stop the processing of the request after the first error.* `Proceed` - Proceed with the processing of the request even when an error occurs. 
* `completion_time`:(string) The timestamp in UTC when the request processing is completed. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_sub_requests`:(int) The number of subrequests received in this request. 
* `request_received_time`:(string) The timestamp in UTC when the request was received. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The processing status of the request.* `NotStarted` - Indicates that the request processing has not begun yet.* `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request.* `ObjPresenceCheckComplete` - Indicates that the object presence check is complete.* `ExecutionInProgress` - Indicates that the request processing is in progress.* `Completed` - Indicates that the request processing has been completed successfully.* `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests.* `Failed` - Indicates that the processing of this request failed.* `TimedOut` - Indicates that the request processing timed out. 
* `status_message`:(string) The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case. 
* `uri`:(string) The URI on which this async operation is being performed. 
 
