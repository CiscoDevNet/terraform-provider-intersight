---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_account"
description: |-
  The Intersight Account used to access Intersight.
---

# Data Source: intersight_iam_account
The Intersight Account used to access Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_account.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Intersight account. By default, name is same as the MoID of the account. 
* `status`:(string) Status of the account. To activate the Intersight account, claim a device to the account. 
 