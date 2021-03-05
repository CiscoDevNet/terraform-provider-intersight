---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_running_firmware"
description: |-
  Running Firmware on an endpoint.
---

# Data Source: intersight_firmware_running_firmware
Running Firmware on an endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_running_firmware.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `component`:(string) Kind of the firmware - boot-booloader/system/kernel. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `package_version`:(string) Package version which the firmware belongs to. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `type`:(string) The type of the firmware. 
* `nr_version`:(string) The version of the firmware. 
 