---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_eula"
description: |-
  End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
---

# Data Source: intersight_firmware_eula
End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_eula.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accepted`:(bool) EULA acceptance status for the account. 
* `content`:(string) EULA acceptance form content provided by cisco.com. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 