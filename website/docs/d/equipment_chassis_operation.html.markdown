---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_chassis_operation"
description: |-
  Models the configurable properties of Chassis.
---

# Data Source: intersight_equipment_chassis_operation
Models the configurable properties of Chassis.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_chassis_operation.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_action`:(string) User configured state of the locator LED for the Chassis.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis. Applying - This state denotes that the settings are being applied in the target chassis. Failed - This state denotes that the settings could not be applied in the target chassis.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 