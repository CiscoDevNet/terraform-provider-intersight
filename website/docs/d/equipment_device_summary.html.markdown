---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_device_summary"
description: |-
  Aggregation of properties pertaining to different inventory MOs.
---

# Data Source: intersight_equipment_device_summary
Aggregation of properties pertaining to different inventory MOs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_device_summary.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string) The distinguished name for the Network Element. 
* `model`:(string) The model information of the Network Element. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `serial`:(string) The serial number for the Network Element. 
* `source_object_type`:(string) The source object type of this view MO. 
 