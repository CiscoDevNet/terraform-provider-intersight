---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_membership_holder"
description: |-
  A holder of REST resources and their membership.
---

# Data Source: intersight_resource_membership_holder
A holder of REST resources and their membership.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_membership_holder.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of this resource membership holder. 
 