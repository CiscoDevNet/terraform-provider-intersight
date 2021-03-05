---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_group"
description: |-
  User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
---

# Data Source: intersight_iam_user_group
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the user group which the dynamic user belongs to. 
 