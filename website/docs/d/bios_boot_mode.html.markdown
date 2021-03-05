---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_boot_mode"
description: |-
  The mode through which bios has booted.
---

# Data Source: intersight_bios_boot_mode
The mode through which bios has booted.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_boot_mode.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `actual_boot_mode`:(string) The actual BIOS boot mode as reported by the platform BIOS. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 