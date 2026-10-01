---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure_disk_slot_ep"
description: |-
        EnclosureDiskSlotEps represent physical disk slots on an enclosure, including slot id, presence, zoning/drive path, and health of the disk in the slot.
        #### Purpose
        Provide per-slot inventory for enclosures to support slot-level presence/health monitoring and zoning visibility.
        #### Key Concepts
        - **Slot inventory:** Identifies each enclosure slot and whether a disk is present.
        - **Zoning/path context:** Captures applied drive path/zoning information.
        - **Health at slot level:** Reports health of the inserted disk as seen from the slot.

---

# Data Source: intersight_storage_enclosure_disk_slot_ep
EnclosureDiskSlotEps represent physical disk slots on an enclosure, including slot id, presence, zoning/drive path, and health of the disk in the slot.
#### Purpose
Provide per-slot inventory for enclosures to support slot-level presence/health monitoring and zoning visibility.
#### Key Concepts
- **Slot inventory:** Identifies each enclosure slot and whether a disk is present.
- **Zoning/path context:** Captures applied drive path/zoning information.
- **Health at slot level:** Reports health of the inserted disk as seen from the slot.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_enclosure_disk_slot_ep.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `drive_path`:(string) This field identifies the zoning configuration applied to  this enclosure slot. 
* `health`:(string) This field identifies the health of the disk inserted in the slot. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field identifies the disk is present in the enclosure slot. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot`:(string) This field represents the slot Id in the storage enclosure. 
 
