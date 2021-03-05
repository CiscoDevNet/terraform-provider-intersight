---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_qualifier"
description: |-
  The qualifier defines how a user qualifies to be part of a user group.
---

# Data Source: intersight_iam_qualifier
The qualifier defines how a user qualifies to be part of a user group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_qualifier.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion. 
 