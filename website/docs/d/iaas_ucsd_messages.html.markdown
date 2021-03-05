---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_messages"
description: |-
  Gets ucsd messages from UCSD.
---

# Data Source: intersight_iaas_ucsd_messages
Gets ucsd messages from UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_ucsd_messages.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `status_id`:(string) Last checked time of the alerts. 
 