
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_identity_summary"
sidebar_current: "docs-intersight-data-source-equipment-identity-summary"
description: |-
Consolidated view of all equipment identities.
---

# Data Source: intersight_equipment._identity_summary
Consolidated view of all equipment identities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `adapter_serial`:(string) Serial Identifier of an adapter participating in SWM. 
* `admin_action`:(string) Updated by UI/API to trigger specific chassis action type.* `None` - No operation value for maintenance actions on an equipment.* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.* `Recommission` - Recommission the equipment.* `Reack` - Reacknowledge the equipment and discover it again.* `Remove` - Remove the equipment permanently from Intersight management. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `chassis_id`:(int) Chassis Identifier of a blade server. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) FI Device registration Mo ID. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. 
* `lifecycle`:(string) The equipment's lifecycle status.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.* `OperationFailed` - Failed Operation Lifecycle state.* `ReackInProgress` - ReackInProgress Lifecycle state.* `RemoveInProgress` - RemoveInProgress Lifecycle state.* `Discovered` - Discovered Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pending_discovery`:(string) Indicates pending discovery flag. 
* `serial`:(string) The serial number of the equipment. 
* `slot_id`:(int) Chassis slot number of a blade server. 
* `source_object_type`:(string) The source object type of this view MO. 
* `switch_id`:(int) Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. 
* `vendor`:(string) The manufacturer of the equipment. 
