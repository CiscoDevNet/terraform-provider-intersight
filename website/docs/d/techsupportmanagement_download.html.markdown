
---
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_download"
sidebar_current: "docs-intersight-data-source-techsupportmanagement-download"
description: |-
Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
---

# Data Source: intersight_techsupportmanagement._download
Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
