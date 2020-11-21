
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_file_downloader"
sidebar_current: "docs-intersight-data-source-niaapi-file-downloader"
description: |-
Provide a presigned url to download the metadata file from server.
---

# Data Source: intersight_niaapi._file_downloader
Provide a presigned url to download the metadata file from server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `file_name`:(string) Filename of this Metadata package file, folder will be handled by api. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `presigned_url`:(string) The presigned URL from server to download this file. 
