---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_diagnostic_messages"
description: |-
  Gets diagnostics messages from UCSD.
---

# Data Source: intersight_iaas_diagnostic_messages
Gets diagnostics messages from UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_diagnostic_messages.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) Message category of the alerts. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `item`:(string) Message target type of the alerts. 
* `last_checked`:(string) Last checked time of the alerts. 
* `message`:(string) Detailed info about the alert. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `recommendation`:(string) Recommended fix for the alert. 
* `status`:(string) Status of the given alert. 
* `status_id`:(string) Status Id of the given alert. 
 