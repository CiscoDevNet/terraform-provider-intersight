---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nia_metadata"
description: |-
  Contains the latest Metadata available for download from server.
---

# Data Source: intersight_niaapi_nia_metadata
Contains the latest Metadata available for download from server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_nia_metadata.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `date`:(string) The date when this package is generated. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the Metadata file downloaded. 
* `metadata_filename`:(string) The Filename of this Metadata package. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(int) The version number of the Metadata package. 
 