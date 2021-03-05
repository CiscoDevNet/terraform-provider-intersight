---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_permission"
description: |-
  Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
---

# Data Source: intersight_iam_permission
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_permission.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) The informative description about each permission. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the permission which has to be granted to user. 
 