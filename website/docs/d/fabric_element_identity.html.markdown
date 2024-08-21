---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_element_identity"
description: |-
        Identity object that uniquely represents a network element object under the domain.

---

# Data Source: intersight_fabric_element_identity
Identity object that uniquely represents a network element object under the domain.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_element_identity.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_action`:(string) Updated by UI/API to trigger specific action type.* `None` - No operation value for maintenance actions on an equipment.* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.* `Recommission` - Recommission the equipment.* `Reack` - Reacknowledge the equipment and discover it again.* `Remove` - Remove the equipment permanently from Intersight management.* `Replace` - Replace the equipment with the other one. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain`:(string) Name of the Fabric Interconnect domain. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. Identifier can only be changed if it has been PATCHED with the AdminAction property set to 'Recommission'. 
* `nr_lifecycle`:(string) The equipment's lifecycle status.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.* `OperationFailed` - Failed Operation Lifecycle state.* `ReackInProgress` - ReackInProgress Lifecycle state.* `RemoveInProgress` - RemoveInProgress Lifecycle state.* `Discovered` - Discovered Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.* `SecureEraseInProgress` - Secure Erase is in progress on given physical entity.* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.* `Inactive` - Inactive Lifecycle state.* `ReplaceInProgress` - ReplaceInProgress Lifecycle state.* `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.* `DomainRmaPendingUserAction` - Domain RMA detected due to the presence of an old pair of FIs with mismatched serial numbers within the same account. User to either initiate the 'Replace Domain workflow' or unclaim the old domain. 
* `lifecycle_mod_time`:(string) The time when the last life cycle state change happened. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the equipment for unique identification. 
* `partial_deployment_status`:(string) Determines if there is partial configuration that has to be deployed on any of the server profiles associated with the server connected to the Fabric Interconnect in cases where one or more server profiles  was deployed when the Fabric Interconnect was down.* `None` - No configuration which is yet to be deployed.The default state of a fabric interconnect which does not have any pending deployment.* `Pending` - There is pending configuration which is yet to be deployed on the fabric interconnect.* `Deploying` - Pending configuration is being deployed on the fabric interconnect. 
* `replacement_type`:(string) Replacement type specifies whether it is single FI or domain replacement.* `None` - The default action is none.* `Individual` - Replacement of single network element.* `Domain` - Domain indicates the replacement of Fabric Interconnect domain. 
* `serial`:(string) The serial number of the equipment. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_id`:(string) Switch Identifier that uniquely represents the fabric object.* `A` - Switch Identifier of Fabric Interconnect A.* `B` - Switch Identifier of Fabric Interconnect B. 
* `vendor`:(string) The manufacturer of the equipment. 
 
