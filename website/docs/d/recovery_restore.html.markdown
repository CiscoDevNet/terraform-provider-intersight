---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_restore"
description: |-
  Triggers a restore operation on the target endpoint.
---

# Data Source: intersight_recovery_restore
Triggers a restore operation on the target endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recovery_restore.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 