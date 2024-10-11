---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_blade_identity"
description: |-
        Identity object that uniquely represents a blade server object under a DR.

---

# Data Source: intersight_compute_blade_identity
Identity object that uniquely represents a blade server object under a DR.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_blade_identity.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_action`:(string) Updated by UI/API to trigger specific action type.* `None` - No operation value for maintenance actions on an equipment.* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.* `Recommission` - Recommission the equipment.* `Reack` - Reacknowledge the equipment and discover it again.* `Remove` - Remove the equipment permanently from Intersight management.* `Replace` - Replace the equipment with the other one. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `chassis_id`:(int) Chassis Identifier of a blade server. 
* `create_time`:(string) The time when this managed object was created. 
* `current_chassis_id`:(int) The id of the chassis that the blade is currently located in. 
* `current_slot_id`:(int) The slot number in the chassis that the blade is currently located in. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_supportability`:(string) Describes whether the running CIMC version supports Intersight managed mode.* `Unknown` - The running firmware version is unknown.* `Supported` - The running firmware version is known and supports IMM mode.* `NotSupported` - The running firmware version is known and does not support IMM mode. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. Identifier can only be changed if it has been PATCHED with the AdminAction property set to 'Recommission'. 
* `last_discovery_triggered`:(string) Denotes the type of discovery that was most recently triggered on this server.* `Unknown` - The last discovery type is unknown.* `Deep` - The last discovery triggered is deep.* `Shallow` - The last discovery triggered is shallow. 
* `nr_lifecycle`:(string) The equipment's lifecycle status.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.* `OperationFailed` - Failed Operation Lifecycle state.* `ReackInProgress` - ReackInProgress Lifecycle state.* `RemoveInProgress` - RemoveInProgress Lifecycle state.* `Discovered` - Discovered Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.* `SecureEraseInProgress` - Secure Erase is in progress on given physical entity.* `ScrubInProgress` - Scrub is in progress on given physical entity.* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.* `Inactive` - Inactive Lifecycle state.* `ReplaceInProgress` - ReplaceInProgress Lifecycle state.* `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.* `ReplacementPendingUserAction` - A new blade server is detected, and discovery is pending cleanup of old blade originally discovered in the new blade's location.* `Removed` - The blade server has been removed from its discovered slot, and not detected anywhere else within the domain.* `Moved` - The blade server has been moved from its discovered location to a new location within the domain.* `Replaced` - The blade server has been removed from its discovered location and another blade has been inserted. in that location.* `MovedAndReplaced` - The blade server has been moved from its discovered location to a new location within the domain and another blade has been inserted into the previously discovered location.* `DomainRmaPendingUserAction` - Domain RMA detected due to the presence of an old pair of FIs with mismatched serial numbers within the same account. User to either initiate the 'Replace Domain workflow' or unclaim the old domain. 
* `lifecycle_mod_time`:(string) The time when the last life cycle state change happened. 
* `manager_slot_id`:(int) Chassis slot number of the manager compute server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the equipment for unique identification. 
* `presence`:(string) The presence state of the blade server.* `Unknown` - The default presence state.* `Equipped` - The server is equipped in the slot.* `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot.* `Missing` - The server is not present in the given slot. 
* `reset_to_default`:(bool) Specifies whether device configurations need to be reset to default upon first-time discovery or recommission of a server. 
* `serial`:(string) The serial number of the equipment. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Chassis slot number of a blade server. 
* `vendor`:(string) The manufacturer of the equipment. 
 
