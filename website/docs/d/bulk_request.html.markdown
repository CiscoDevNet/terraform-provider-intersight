---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_request"
description: |-
        Requests (bulk.Request) represent a bulk-operation submission that allows clients to execute multiple API actions—such as Create, Update, or Delete—against a given base URI, including operations across multiple subpaths relative to that URI. The object is designed for high-volume changes with optional asynchronous response handling.
        #### Purpose
        Enable efficient, controlled execution of many related API changes in a single bulk request, reducing client-side overhead and supporting consistent execution semantics for batch create/update/delete workflows.
        #### Key Concepts
        - **Bulk API execution container**: Encapsulates a set of actions (`actions`) that the platform will apply in bulk relative to a provided target URI.
        - **Multi-subpath operations**: Supports acting on multiple objects/resources under the same REST resource type (e.g., PATCH multiple objects of a type).
        - **Action semantics and pre-checks**: `actions` describe what should be done, and can include behavior such as checking for existence vs executing changes (as described in the model).
        - **Asynchronous support**: CREATE supports `respond-async: true`, allowing clients to request async processing suitable for large batches.
        - **Create-only action definition**: `actions` is `createonly`, preserving the integrity of what was submitted once the bulk request is created.
        - **Global privilege gating**: Access is controlled via global privilege sets for READ/CREATE, reflecting the broad impact and cross-resource nature of bulk operations.
        - **Organization-owned request record**: `owner: organization` indicates the bulk request exists within an org/account context even though the privileges are global.

---

# Data Source: intersight_bulk_request
Requests (bulk.Request) represent a bulk-operation submission that allows clients to execute multiple API actions—such as Create, Update, or Delete—against a given base URI, including operations across multiple subpaths relative to that URI. The object is designed for high-volume changes with optional asynchronous response handling.
#### Purpose
Enable efficient, controlled execution of many related API changes in a single bulk request, reducing client-side overhead and supporting consistent execution semantics for batch create/update/delete workflows.
#### Key Concepts
- **Bulk API execution container**: Encapsulates a set of actions (`actions`) that the platform will apply in bulk relative to a provided target URI.
- **Multi-subpath operations**: Supports acting on multiple objects/resources under the same REST resource type (e.g., PATCH multiple objects of a type).
- **Action semantics and pre-checks**: `actions` describe what should be done, and can include behavior such as checking for existence vs executing changes (as described in the model).
- **Asynchronous support**: CREATE supports `respond-async: true`, allowing clients to request async processing suitable for large batches.
- **Create-only action definition**: `actions` is `createonly`, preserving the integrity of what was submitted once the bulk request is created.
- **Global privilege gating**: Access is controlled via global privilege sets for READ/CREATE, reflecting the broad impact and cross-resource nature of bulk operations.
- **Organization-owned request record**: `owner: organization` indicates the bulk request exists within an org/account context even though the privileges are global.
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
* `is_backup_encryption_key_set`:(bool) Indicates whether the value of the 'backupEncryptionKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_sub_requests`:(int) The number of sub requests received in this request. 
* `org_moid`:(string) The moid of the organization under which this request was issued. 
* `request_received_time`:(string) The timestamp when the request was received. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_duplicates`:(bool) Skip the already present objects. 
* `status`:(string) The processing status of the Request.* `NotStarted` - Indicates that the request processing has not begun yet.* `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request.* `ObjPresenceCheckComplete` - Indicates that the object presence check is complete.* `ExecutionInProgress` - Indicates that the request processing is in progress.* `Completed` - Indicates that the request processing has been completed successfully.* `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests.* `Failed` - Indicates that the processing of this request failed.* `TimedOut` - Indicates that the request processing timed out. 
* `status_message`:(string) The status message corresponding to the status. 
* `uri`:(string) The URI on which this bulk action is to be performed.The value will be used when there is no override in the SubRequest. 
* `verb`:(string) The type of operation to be performed.One of - Post (Create), Patch (Update) or Delete (Remove).The value will be used when there is no override in the SubRequest.* `POST` - Used to create a REST resource.* `PATCH` - Used to update a REST resource.* `DELETE` - Used to delete a REST resource. 
 
