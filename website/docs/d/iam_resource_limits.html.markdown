---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_limits"
description: |-
  The resource limits used to limit resources such as User accounts.
---

# Data Source: intersight_iam_resource_limits
The resource limits used to limit resources such as User accounts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_resource_limits.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `per_account_user_limit`:(int) The maximum number of users allowed in an account. The default value is 200. 
 