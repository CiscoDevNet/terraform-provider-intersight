---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege_set"
description: |-
  A collection of privileges.
---

# Data Source: intersight_iam_privilege_set
A collection of privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_privilege_set.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the privilege set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the privilege set. 
 