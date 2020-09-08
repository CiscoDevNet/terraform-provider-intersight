
---
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_tech_support_status"
sidebar_current: "docs-intersight-data-source-techsupportmanagement-tech-support-status"
description: |-
The techsupport collection status.
---

# Data Source: intersight_techsupportmanagement._tech_support_status
The techsupport collection status.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `file_name`:(string) The name of the Techsupport bundle file. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `reason`:(string) Reason for techsupport failure, if any. 
* `relay_reason`:(string) Reason for status relay failure, if any. 
* `relay_status`:(string) Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. 
* `request_ts`:(string) The time at which the techsupport request was initiated. 
* `status`:(string) Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. 
* `techsupport_download_url`:(string) The Url to download the techsupport file. 
