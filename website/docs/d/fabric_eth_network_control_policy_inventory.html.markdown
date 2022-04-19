---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_control_policy_inventory"
description: |-
        The features that are applied on a vethernet that is connected to the vNIC.

---

# Data Source: intersight_fabric_eth_network_control_policy_inventory
The features that are applied on a vethernet that is connected to the vNIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_eth_network_control_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cdp_enabled`:(bool) Enables the CDP on an interface. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `forge_mac`:(string) Determines if the MAC forging is allowed or denied on an interface.* `allow` - Allows mac forging on an interface.* `deny` - Denies mac forging on an interface. 
* `mac_registration_mode`:(string) Determines the MAC addresses that have to be registered with the switch.* `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN.* `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uplink_fail_action`:(string) Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned.* `linkDown` - The vethernet will go down in case a suitable uplink is not pinned.* `warning` - The vethernet will remain up even if a suitable uplink is not pinned. 
 
