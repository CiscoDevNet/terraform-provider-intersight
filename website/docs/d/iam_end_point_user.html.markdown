---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user"
description: |-
  Endpoint User or Local User.
---

# Data Source: intersight_iam_end_point_user
Endpoint User or Local User.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters. 
 