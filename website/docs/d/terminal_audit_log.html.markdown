---
subcategory: "terminal"
layout: "intersight"
page_title: "Intersight: intersight_terminal_audit_log"
description: |-
  Audit log of remote terminal user sessions.
---

# Data Source: intersight_terminal_audit_log
Audit log of remote terminal user sessions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_terminal_audit_log.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `end_time`:(string) The time the terminal was closed. If terminal has not closed, value is zero time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `start_time`:(string) The time the terminal session was opened. 
 