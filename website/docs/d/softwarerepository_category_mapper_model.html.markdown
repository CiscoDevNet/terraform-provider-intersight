
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper_model"
sidebar_current: "docs-intersight-data-source-softwarerepository-category-mapper-model"
description: |-
Maps a Cisco hardware model Series to its applicable hardware models.
---

# Data Source: intersight_softwarerepository._category_mapper_model
Maps a Cisco hardware model Series to its applicable hardware models.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) The category of the model series. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dist_tag`:(string) The distributable tag value of the model series. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `regex_pattern`:(string) The regex that all images of this model follow. 
* `series_id`:(string) Cisco hardware model series. 
