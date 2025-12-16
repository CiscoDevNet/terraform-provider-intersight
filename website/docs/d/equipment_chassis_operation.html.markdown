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
To access the ith object of the results obtained, use `data.intersight_equipment_chassis_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_locator_led_action`:(string) User configured state of the locator LED for the Chassis.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `admin_power_cycle_expander_module_slot_id`:(int) Slot id of the expander module slot within chassis that needs to be power cycled. 
* `admin_power_cycle_slot_id`:(int) Slot id of the chassis slot that needs to be power cycled. 
* `admin_reset_config_slot_id`:(int) Slot id of the chassis slot that needs to have its configuration reset. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis. Applying - This state denotes that the settings are being applied in the target chassis. Failed - This state denotes that the settings could not be applied in the target chassis.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
