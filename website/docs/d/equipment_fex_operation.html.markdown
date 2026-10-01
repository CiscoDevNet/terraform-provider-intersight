---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex_operation"
description: |-
        The FexOperation object models operational actions that can be executed on a Fabric Extender (FEX), such as locator LED control, with workflow-backed status reporting.
        #### Purpose
        FexOperation provides a consistent API mechanism to initiate FEX maintenance actions and observe their completion state, enabling safe and auditable operational control of fabric extenders.
        ### Key Concepts
        - **FEX operational actions:** Represents targeted maintenance actions for a specific FEX instance.
        - **Workflow-driven execution:** Tracks operation progress and completion via action/config state.
        - **Inventory correlation:** Tied to the inventoried FEX object to ensure actions target the correct hardware.
        - **Access control and safety:** Uses privilege gating for operational actions to minimize disruption risk.

---

# Data Source: intersight_equipment_fex_operation
The FexOperation object models operational actions that can be executed on a Fabric Extender (FEX), such as locator LED control, with workflow-backed status reporting.
#### Purpose
FexOperation provides a consistent API mechanism to initiate FEX maintenance actions and observe their completion state, enabling safe and auditable operational control of fabric extenders.
### Key Concepts
- **FEX operational actions:** Represents targeted maintenance actions for a specific FEX instance.
- **Workflow-driven execution:** Tracks operation progress and completion via action/config state.
- **Inventory correlation:** Tied to the inventoried FEX object to ensure actions target the correct hardware.
- **Access control and safety:** Uses privilege gating for operational actions to minimize disruption risk.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_fex_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_locator_led_action`:(string) Action performed on the locator LED for a FEX.* `None` - No operation action for the Locator Led of an equipment.* `TurnOn` - Turn on the Locator Led of an equipment.* `TurnOff` - Turn off the Locator Led of an equipment. 
* `admin_locator_led_action_state`:(string) Defines status of action performed on AdminLocatorLedState.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
