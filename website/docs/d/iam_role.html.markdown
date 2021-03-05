---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_role"
description: |-
  A role is a collection of privilege sets that are assigned to a user using a permission object.
---

# Data Source: intersight_iam_role
A role is a collection of privilege sets that are assigned to a user using a permission object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_role.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Informative description about each role. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the role which has to be granted to user. 
 