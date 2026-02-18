---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_rack_enclosure_slot"
description: |-
        The equipment.RackEnclosureSlot object represents a physical slot within a rack enclosure that is designed to hold a rack-mounted server.
        #### Purpose
        The main function of this object is to model the individual server slots within a larger rack enclosure. It establishes the link between a physical slot and the compute.RackUnit that occupies it. It is essential for accurately representing the physical topology of the enclosure and its contained servers.
        #### Key Concepts
        - **Physical Slot Representation:** Models a single server bay within a rack enclosure.
        - **Server-to-Slot Mapping:** The relationship to a compute.RackUnit and the rackUnitDn property create the explicit link between the server and its physical location.
        - **Location Identification:** The rackId property (equivalent to the server ID) and its parent equipment.RackEnclosure uniquely identify the slot's position.
        - **Hierarchical Structure:** Exists as a child of an equipment.RackEnclosure, forming part of the overall enclosure inventory.

---

# Data Source: intersight_equipment_rack_enclosure_slot
The equipment.RackEnclosureSlot object represents a physical slot within a rack enclosure that is designed to hold a rack-mounted server.
#### Purpose
The main function of this object is to model the individual server slots within a larger rack enclosure. It establishes the link between a physical slot and the compute.RackUnit that occupies it. It is essential for accurately representing the physical topology of the enclosure and its contained servers.
#### Key Concepts
- **Physical Slot Representation:** Models a single server bay within a rack enclosure.
- **Server-to-Slot Mapping:** The relationship to a compute.RackUnit and the rackUnitDn property create the explicit link between the server and its physical location.
- **Location Identification:** The rackId property (equivalent to the server ID) and its parent equipment.RackEnclosure uniquely identify the slot's position.
- **Hierarchical Structure:** Exists as a child of an equipment.RackEnclosure, forming part of the overall enclosure inventory.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_rack_enclosure_slot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `rack_id`:(int) Server ID which is part of Rack Enclosure Slot. 
* `rack_unit_dn`:(string) Server DN which is part of Rack Enclosure Slot. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
