---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_slot_array"
description: |-
        Type to represent additional switch specific capabilities.

---

# Data Source: intersight_capability_equipment_slot_array
Type to represent additional switch specific capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_equipment_slot_array.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `first_index`:(float) First Index information for a Switch/Fabric-Interconnect hardware. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect hardware. 
* `horizontal_start_offset`:(float) Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_separation`:(float) Inline Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `inline_group_size`:(float) Inline Group Size information for a Switch/Fabric-Interconnect hardware. 
* `inline_offset`:(float) Inline Offset information for a Switch/Fabric-Interconnect hardware. 
* `location`:(string) Location information for a Switch/Fabric-Interconnect hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `number_of_slots`:(int) Number of Slots information for a Switch/Fabric-Interconnect hardware. 
* `orientation`:(string) Orientation information for a Switch/Fabric-Interconnect hardware. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `UCS-FI-6536` - The standard 5th generation UCS Fabric Interconnect with 36 ports.* `UCSX-S9108-100G` - Cisco UCS Fabric Interconnect 9108 100G with 8 ports.* `UCS-FI-6664` - The standard 6th generation UCS Fabric Interconnect with 36 ports.* `unknown` - Unknown device type, usage is TBD. 
* `selector`:(string) Selector information for a Switch/Fabric-Interconnect hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `slots_per_line`:(int) Slots per line information for a Switch/Fabric-Interconnect hardware. 
* `transverse_group_separation`:(float) Transverse Group Separation information for a Switch/Fabric-Interconnect hardware. 
* `transverse_group_size`:(float) Transverse Group Size information for a Switch/Fabric-Interconnect hardware. 
* `transverse_offset`:(float) Transverse Offset information for a Switch/Fabric-Interconnect hardware. 
* `vertical_start_offset`:(float) Vertical Start Offset information for a Switch/Fabric-Interconnect hardware. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect hardware. 
 
