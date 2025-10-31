---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_node"
description: |-
        External PCIe nodes connected to a server.

---

# Data Source: intersight_pci_node
External PCIe nodes connected to a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chassis_id`:(string) The id of the chassis that the PCIe node is currently located in. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) The version of the PCIe node CIMC firmware. 
* `inventory_ready`:(bool) The inventory ready field indicates whether the PCIe node management controller has completed inventory of the installed PCIe devices. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `nr_lifecycle`:(string) The lifecycle state of the PCIe node. This will map to the discovery lifecycle as represented in the Identity object.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.* `SecureEraseInProgress` - Secure Erase is in progress on given physical entity.* `ScrubInProgress` - Scrub is in progress on given physical entity.* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.* `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.* `Removed` - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade.* `Moved` - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade.* `Replaced` - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade.* `MovedAndReplaced` - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the PCIe node, the value of this property is the name of the UCS Domain along with the chassis and node slot ids. 
* `oper_state`:(string) Operational state of the PCIe node. 
* `package_version`:(string) Bundle version which the CIMC firmware belongs to. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(string) The slot number in the chassis that the PCIe node is currently located in. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
