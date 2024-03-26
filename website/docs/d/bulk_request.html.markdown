---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_request"
description: |-
        The bulk.Request API allows users to perform API actions (Create, Update or Delete) in bulk, on a given URI.
        It is possible to operate on multiple subpaths relative to the provided URI (For example, it would be possible to
        perform a PATCH action on multiple objects of a given REST resource type).

---

# Data Source: intersight_bulk_request
The bulk.Request API allows users to perform API actions (Create, Update or Delete) in bulk, on a given URI.
It is possible to operate on multiple subpaths relative to the provided URI (For example, it would be possible to
perform a PATCH action on multiple objects of a given REST resource type).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action_on_error`:(string) The action to be taken when an error occurs during processing of the request.* `Stop` - Stop the processing of the request after the first error.* `Proceed` - Proceed with the processing of the request even when an error occurs. 
* `completion_time`:(string) The timestamp when the request processing completed. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_sub_requests`:(int) The number of sub requests received in this request. 
* `org_moid`:(string) The moid of the organization under which this request was issued. 
* `request_received_time`:(string) The timestamp when the request was received. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_duplicates`:(bool) Skip the already present objects. 
* `status`:(string) The processing status of the Request.* `NotStarted` - Indicates that the request processing has not begun yet.* `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request.* `ObjPresenceCheckComplete` - Indicates that the object presence check is complete.* `ExecutionInProgress` - Indicates that the request processing is in progress.* `Completed` - Indicates that the request processing has been completed successfully.* `Failed` - Indicates that the processing of this request failed.* `TimedOut` - Indicates that the request processing timed out. 
* `status_message`:(string) The status message corresponding to the status. 
* `uri`:(string) The URI on which this bulk action is to be performed.The value will be used when there is no override in the SubRequest. 
* `verb`:(string) The type of operation to be performed.One of - Post (Create), Patch (Update) or Delete (Remove).The value will be used when there is no override in the SubRequest.* `POST` - Used to create a REST resource.* `PATCH` - Used to update a REST resource.* `DELETE` - Used to delete a REST resource. 
 
