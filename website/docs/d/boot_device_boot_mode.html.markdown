---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_device_boot_mode"
description: |-
  Boot mode of the devices that BIOS uses to boot them.
---

# Data Source: intersight_boot_device_boot_mode
Boot mode of the devices that BIOS uses to boot them.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_device_boot_mode.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `configured_boot_mode`:(string) The user desired BIOS boot mode as configured in the boot policy. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 