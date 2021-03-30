---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_auto_rma_policy"
description: |-
  Auto rma policy to decide whether rma data should be collected
---

# Data Source: intersight_appliance_auto_rma_policy
Auto rma policy to decide whether rma data should be collected
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_auto_rma_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enable`:(bool) Status of the data collection mode. If the value is 'true', then data collection is enabled. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 