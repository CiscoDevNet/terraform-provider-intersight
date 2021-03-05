---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper_model"
description: |-
  Maps a Cisco hardware model Series to its applicable hardware models.
---

# Data Source: intersight_softwarerepository_category_mapper_model
Maps a Cisco hardware model Series to its applicable hardware models.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_softwarerepository_category_mapper_model.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) The category of the model series. 
* `dist_tag`:(string) The distributable tag value of the model series. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `regex_pattern`:(string) The regex that all images of this model follow. 
* `series_id`:(string) Cisco hardware model series. 
 