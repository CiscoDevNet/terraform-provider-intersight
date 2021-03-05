---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_capability_def"
description: |-
  Chassis Iocard module capabilities.
---

# Data Source: intersight_capability_io_card_capability_def
Chassis Iocard module capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_io_card_capability_def.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dc_supported`:(bool) Device connector support on Iocard. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
 