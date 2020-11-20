
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex_operation"
sidebar_current: "docs-intersight-data-source-equipment-fex-operation"
description: |-
Models the configuration states of a FEX in Intersight.
---

# Data Source: intersight_equipment._fex_operation
Models the configuration states of a FEX in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_action`:(string) Action performed on the locator LED for a FEX.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `admin_locator_led_action_state`:(string) Defines status of action performed on AdminLocatorLedState.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
