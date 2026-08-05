---
subcategory: "power"
layout: "intersight"
page_title: "Intersight: intersight_power_power_group_member"
description: |-
        Managed object used to track power information for a member of a power group.

---

# Data Source: intersight_power_power_group_member
Managed object used to track power information for a member of a power group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_power_power_group_member.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocated_power`:(int) Power allocated to the power group member by the power group in watts. 
* `config_state`:(string) Configuration State of the power group member.* `Ok` - The configuration state of the power group member is ok.* `Unknown` - The configuration state of the power group member is not known.* `PendingRemoval` - The power group member is pending removal. The member has been removed from the power group's InventoryDeviceMember list but the workflow to remove power limit from the device has not yet completed. Once the workflow succeeds, this member MO will be deleted. 
* `create_time`:(string) The time when this managed object was created. 
* `device_registration_moid`:(string) Moid of the asset.DeviceRegistration corresponding to this member's inventory device. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `max_required_power`:(int) Maximum power required by the power group member in watts. 
* `member_type`:(string) The type of the power group member - Chassis, Fex or Fabric Interconnect.* `Unknown` - The power group member type is unknown.* `Chassis` - The power group member is a chassis.* `Fex` - The power group member is a Fex.* `FabricInterconnect` - The power group member is a Fabric Interconnect. 
* `min_required_power`:(int) Minimum power required by the power group member in watts. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational State of the power group member.* `Ok` - The operational state of the power group member is ok.* `PowerConsumptionNearingLimit` - Operational state for power group member is PowerConsumptionNearingLimit. This may happen when the current power consumption of a chassis approaches the chassis power limit. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
