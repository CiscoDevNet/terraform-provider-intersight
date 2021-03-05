---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_group_status"
description: |-
  Status of a group of applications.
---

# Data Source: intersight_appliance_group_status
Status of a group of applications.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_group_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the group. 
* `group_name`:(string) The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `overall_status`:(string) The overall API status from this group's applications. 
 