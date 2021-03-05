---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_catalog"
description: |-
  A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
---

# Data Source: intersight_softwarerepository_catalog
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_image_pull_failure`:(bool) The status of the image catalog synchronization operation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the catalog. The names are populated and predefined during MO creation. 
 