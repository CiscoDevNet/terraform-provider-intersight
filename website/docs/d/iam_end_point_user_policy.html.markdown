---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_policy"
description: |-
  Enables creation of local users on endpoints.
---

# Data Source: intersight_iam_end_point_user_policy
Enables creation of local users on endpoints.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 