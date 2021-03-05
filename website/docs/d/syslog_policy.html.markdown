---
subcategory: "syslog"
layout: "intersight"
page_title: "Intersight: intersight_syslog_policy"
description: |-
  The syslog policy configure the syslog server to receive CIMC log entries.
---

# Data Source: intersight_syslog_policy
The syslog policy configure the syslog server to receive CIMC log entries.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_syslog_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 