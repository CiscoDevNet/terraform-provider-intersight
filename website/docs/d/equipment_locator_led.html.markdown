---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_locator_led"
description: |-
  Locator Led of an Equipment like Rack, Disk etc.
---

# Data Source: intersight_equipment_locator_led
Locator Led of an Equipment like Rack, Disk etc.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_locator_led.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `color`:(string) Color of the locatorled available on an equipment. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Identifies the operational state of locatorled. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 