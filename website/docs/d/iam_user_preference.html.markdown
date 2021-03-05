---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_user_preference"
description: |-
  Holder for UI preferences such as theme, columns.
---

# Data Source: intersight_iam_user_preference
Holder for UI preferences such as theme, columns.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_user_preference.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `user_unique_identifier`:(string) Unique id of the user used by the identity provider to store the user. 
 