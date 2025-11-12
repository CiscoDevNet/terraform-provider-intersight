---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_switch_operation"
description: |-
        Models the operational states of a Switch in Intersight.

---

# Data Source: intersight_equipment_switch_operation
Models the operational states of a Switch in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_switch_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_evac_state`:(string) Sets evacuation state of the switch. When evacuation is enabled, data traffic flowing through this switch will be suspended for all the servers. Fabric evacuation can be enabled during any maintenance activity on the switch in order to gracefully failover data flows to the peer switch.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `admin_locator_led_action`:(string) Action performed on the locator LED of the switch.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `admin_locator_led_action_state`:(string) Defines status of action performed on AdminLocatorLedState.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `config_evac_state`:(string) Captures the status of evacuation on this switch.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `force_evac`:(bool) Evacuation is blocked by the system if it can cause a traffic outage in the domain. Select \ Force Evacuation\  only if system rejects the operation and you want to override that. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reset_action_state`:(string) Current status of the reset operation executed on the Fabric Interconnect.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_name`:(string) Name of the switch on which the switch operation is performed. 
 
