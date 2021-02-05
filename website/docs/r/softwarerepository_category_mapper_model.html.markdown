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
* `category`:(string) The category of the model series. 
* `dist_tag`:(string) The distributable tag value of the model series. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `regex_pattern`:(string) The regex that all images of this model follow. 
* `series_id`:(string) Cisco hardware model series. 
* `supported_models`:
                (Array of schema.TypeString) -
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_softwarerepository_category_mapper_model` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_softwarerepository_category_mapper_model.example 1234567890987654321abcde
``` 