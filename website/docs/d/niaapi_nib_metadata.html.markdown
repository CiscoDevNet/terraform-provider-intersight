---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nib_metadata"
description: |-
  Contains the latest metadata available for download from server.
---

# Data Source: intersight_niaapi_nib_metadata
Contains the latest metadata available for download from server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_nib_metadata.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `date`:(string) The date when the package was generated. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the metadata file downloaded. 
* `metadata_filename`:(string) The filename of the metadata package. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(int) The version number of the metadata package. 
 