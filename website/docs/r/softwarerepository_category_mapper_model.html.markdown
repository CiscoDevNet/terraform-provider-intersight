---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper_model"
description: |-
  Maps a Cisco hardware model Series to its applicable hardware models.
---

# Resource: intersight_softwarerepository_category_mapper_model
Maps a Cisco hardware model Series to its applicable hardware models.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `category`:(string) The category of the model series. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `dist_tag`:(string) The distributable tag value of the model series. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `regex_pattern`:(string) The regex that all images of this model follow. 
* `series_id`:(string) Cisco hardware model series. 
* `supported_models`:
                (Array of schema.TypeString) -
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_softwarerepository_category_mapper_model` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_softwarerepository_category_mapper_model.example 1234567890987654321abcde
```