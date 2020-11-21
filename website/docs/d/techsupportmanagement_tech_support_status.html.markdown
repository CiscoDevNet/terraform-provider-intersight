
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `file_name`:(string) The name of the Techsupport bundle file. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `reason`:(string) Reason for techsupport failure, if any. 
* `relay_reason`:(string) Reason for status relay failure, if any. 
* `relay_status`:(string) Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. 
* `request_ts`:(string) The time at which the techsupport request was initiated. 
* `status`:(string) Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. 
* `techsupport_download_url`:(string) The Url to download the techsupport file. 
