---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex_operation"
description: |-
  Models the configuration states of a FEX in Intersight.
---

# Data Source: intersight_equipment_fex_operation
Models the configuration states of a FEX in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_fex_operation.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_action`:(string) Action performed on the locator LED for a FEX.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `admin_locator_led_action_state`:(string) Defines status of action performed on AdminLocatorLedState.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 