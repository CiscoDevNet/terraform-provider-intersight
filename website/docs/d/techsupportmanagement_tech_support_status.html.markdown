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
To access the ith object of the results obtained, use `data.intersight_techsupportmanagement_tech_support_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `file_name`:(string) The name of the Techsupport bundle file. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reason`:(string) Reason for techsupport failure, if any. 
* `relay_reason`:(string) Reason for status relay failure, if any. 
* `relay_status`:(string) Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. 
* `request_ts`:(string) The time at which the techsupport request was initiated. 
* `status`:(string) Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. 
* `techsupport_download_url`:(string) The Url to download the techsupport file. 
 