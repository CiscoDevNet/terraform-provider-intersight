---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_operation"
description: |-
        The PcOperation object represents an operational control surface for port-channels, allowing administrative operations such as enabling/disabling or related admin-driven actions.
        #### Purpose
        PcOperation provides an API-facing model for port-channel operational actions that affect the current state of a port-channel on the switch. It supports controlled, auditable port-channel state changes through managed workflows.
        #### Key Concepts
        - **Operational action model:** Represents “do” operations on a port-channel rather than desired steady-state design intent.
        - **Administrative state control:** Enables controlled enable/disable-like operations via API.
        - **Workflow-backed behavior:** Often implemented through asynchronous workflows and config-state reporting.
        - **Network-element scoping:** Actions are tied to a specific network element context.

---

# Data Source: intersight_fabric_pc_operation
The PcOperation object represents an operational control surface for port-channels, allowing administrative operations such as enabling/disabling or related admin-driven actions.
#### Purpose
PcOperation provides an API-facing model for port-channel operational actions that affect the current state of a port-channel on the switch. It supports controlled, auditable port-channel state changes through managed workflows.
#### Key Concepts
- **Operational action model:** Represents “do” operations on a port-channel rather than desired steady-state design intent.
- **Administrative state control:** Enables controlled enable/disable-like operations via API.
- **Workflow-backed behavior:** Often implemented through asynchronous workflows and config-state reporting.
- **Network-element scoping:** Actions are tied to a specific network element context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_pc_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_action`:(string) An operation that has to be perfomed on the port channel. Default value is None which means there will be no implicit port operation triggered.* `None` - No admin triggered action.* `SetUserLabel` - Admin triggered operation to set the user label on the port channel. 
* `admin_state`:(string) Admin configured state to disable the port channel.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_label`:(string) The user defined label assigned to the a Port. 
 
