---
subcategory: "power"
layout: "intersight"
page_title: "Intersight: intersight_power_power_group"
description: |-
        A Power Group represents a collection of equipment that share the same power budget.

---

# Data Source: intersight_power_power_group
A Power Group represents a collection of equipment that share the same power budget.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_power_power_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Administrative state of the power group. Determines whether the power group is enabled or disabled.* `Enabled` - The power group is administratively enabled.* `Disabled` - The power group is administratively disabled. 
* `buffer_power_required`:(int) Reserved power buffer calculated as 10% of total chassis minimum power requirements. FI and FEX do not contribute to buffer power. 
* `config_state`:(string) Power Group Configuration State.* `Ok` - The configuration state of the power group is ok. All members are configured successfully.* `Unknown` - The configuration state of the power group is not known.* `Configuring` - The power group is actively being configured. Transitional state during CREATE or UPDATE operations.* `PartiallyConfigured` - The power group configuration failed mid-way. Some members may be configured while others are not. Totals may not match actual membership. Reconciliation or manual intervention is required.* `PartiallyDeleted` - The configuration state of the power group is Partially Deleted. This occurs when a DELETE operation was initiated but the workflow to remove power limits did not succeed for all members.* `Disabling` - The power group is being disabled. A workflow is in progress to push 0 power limits to all member devices.* `PartiallyDisabled` - The power group disable operation failed mid-way. Some members may have their power limits removed while others retain their previous values. The power group remains in a disabled state but may require manual intervention. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the power group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User defined name for the power group. 
* `oper_state`:(string) Power Group Operational State.* `Ok` - The operational state of the power group is ok. All members are configured successfully.* `Unknown` - The operational state of the power group is not known.* `InsufficientPower` - The power group does not have sufficient power budget to meet the minimum power requirements of all its chassis members. This typically occurs when new hardware is added to a member chassis, increasing its minimum power requirement beyond the available budget. 
* `power_budget`:(int) Maximum power budget allocated to this power group in watts. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_maximum_power_required`:(int) Sum of maximum power required for each member in this group in watts. 
* `total_minimum_power_required`:(int) Sum of minimum power required for each member in this group in watts. 
 
