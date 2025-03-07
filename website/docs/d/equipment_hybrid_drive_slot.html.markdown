---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_hybrid_drive_slot"
description: |-
        NVMe HybridDriveSlots present in a server.

---

# Data Source: intersight_equipment_hybrid_drive_slot
NVMe HybridDriveSlots present in a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_hybrid_drive_slot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `current_mode`:(string) The Configured Mode of the Hybrid Drive slot.* `` - Hybrid Drive slot mode is not applicable.* `Controller` - Hybrid Drive slot mode is Controller.* `Direct` - Hybrid Drive slot mode is Direct. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `requested_mode`:(string) The Requested Mode for the Hybrid Drive slot.* `` - Hybrid Drive slot mode is not applicable.* `Controller` - Hybrid Drive slot mode is Controller.* `Direct` - Hybrid Drive slot mode is Direct. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(string) This field identifies a specific slot on the server. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
