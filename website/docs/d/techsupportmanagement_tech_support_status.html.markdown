---
subcategory: "techsupportmanagement"
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_tech_support_status"
description: |-
        The techsupport collection status.

---

# Data Source: intersight_techsupportmanagement_tech_support_status
The techsupport collection status.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_techsupportmanagement_tech_support_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) The name of the Techsupport bundle file. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reason`:(string) Reason for techsupport failure, if any. 
* `relay_reason`:(string) Reason for status relay failure, if any. 
* `relay_status`:(string) Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. 
* `request_ts`:(string) The time at which the techsupport request was initiated. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. 
* `techsupport_download_url`:(string) The Url to download the techsupport file. 
 
