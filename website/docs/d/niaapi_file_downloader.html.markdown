---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_file_downloader"
description: |-
  Provide a presigned url to download the metadata file from server.
---

# Data Source: intersight_niaapi_file_downloader
Provide a presigned url to download the metadata file from server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_file_downloader.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `file_name`:(string) Filename of this Metadata package file, folder will be handled by api. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presigned_url`:(string) The presigned URL from server to download this file. 
 