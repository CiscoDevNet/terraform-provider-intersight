
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `dist_tag`:(string) The distributable tag value of the model series. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `regex_pattern`:(string) The regex that all images of this model follow. 
* `series_id`:(string) Cisco hardware model series. 
