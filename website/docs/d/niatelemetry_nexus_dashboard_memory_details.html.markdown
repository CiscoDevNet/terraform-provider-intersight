---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_dashboard_memory_details"
description: |-
  Details of Nexus Dashboard's memory.
---

# Data Source: intersight_niatelemetry_nexus_dashboard_memory_details
Details of Nexus Dashboard's memory.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboard_memory_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_name`:(string) Name of the node in Nexus Dashboard cluster. 
* `memory_capacity`:(int) Memory capacity of a node in Nexus Dashboard. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 