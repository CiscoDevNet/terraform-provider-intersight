---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_mapping"
description: |-
  Virtual Media image uploaded on the server.
---

# Data Source: intersight_compute_mapping
Virtual Media image uploaded on the server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_mapping.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `file_location`:(string) Remote image location from where the image is uploaded to server. 
* `identifier`:(string) The identity assigned to this Virtual Media Image by server. 
* `image_name`:(string) Image name of uploaded Virtual Media Image. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of Virtual Media mapping assigne by server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 