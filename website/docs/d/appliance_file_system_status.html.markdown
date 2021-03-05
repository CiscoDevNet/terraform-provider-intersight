---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_system_status"
description: |-
  Status of a file system on an Intersight Appliance node.
---

# Data Source: intersight_appliance_file_system_status
Status of a file system on an Intersight Appliance node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_file_system_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(int) Capacity of the file system in bytes. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mountpoint`:(string) Mount point of this file system. 
* `operational_status`:(string) Operational status of the file system.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
* `usage`:(float) Percentage of the file system capacity currently in use. 
 