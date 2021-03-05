---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_external_syslog_setting"
description: |-
  Configure External Syslog Server.
---

# Data Source: intersight_appliance_external_syslog_setting
Configure External Syslog Server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_external_syslog_setting.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enabled`:(bool) Enable or disable External Syslog Server. 
* `export_nginx`:(bool) Enable or disable exporting of Web Server access logs. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port`:(int) External Syslog Server Port for connection establishment. 
* `server`:(string) External Syslog Server Address, can be IP address or hostname. 
 