---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_count"
description: |-
  Total number of advisories currently affecting a given Account.
---

# Data Source: intersight_tam_advisory_count
Total number of advisories currently affecting a given Account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_count.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_count`:(int) Total number of advisories affecting the account. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 