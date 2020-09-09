
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_download_spec"
sidebar_current: "docs-intersight-data-source-softwarerepository-download-spec"
description: |-
The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
---

# Data Source: intersight_softwarerepository._download_spec
The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `auth_token`:(string) The OAuth2 token that will be used during image download by the endpoint to authenticate with file server. 
* `certificate`:(string) The certificate of file server that will be used by the endpoint to validate the server before starting the file download. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `filename`:(string) The name of the firmware image. 
* `md5sum`:(string) MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `size`:(int) The size (in bytes) of the firmware image. 
* `url`:(string) The URL of this file in file server. The endpoint uses this URL to download the file from the file server. 
