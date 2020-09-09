
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
