---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_dashboard_controller_details"
description: |-
  Details of controller added to NexusDashboard.
---

# Data Source: intersight_niatelemetry_nexus_dashboard_controller_details
Details of controller added to NexusDashboard.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboard_controller_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `site_health`:(int) Health of the site serviced by ND. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `version_of_controller`:(string) Version of the controller serviced by ND. 
 