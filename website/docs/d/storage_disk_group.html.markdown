---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group"
description: |-
        DiskGroups represent a group of one or more Spans used to configure virtual drives. They capture disk-group identity and RAID intent and tie together spans, virtual drives, and dedicated hot spares under a specific storage controller.
        #### Purpose
        Provide a controller-scoped construct for organizing disks (via spans) into RAID-oriented groups that can then produce one or more virtual drives.
        #### Key Concepts
        - **RAID grouping construct**: `raidType` expresses the RAID level intended for virtual drives created from the disk group.
        - **Human-identifiable naming**: `name` identifies the disk group on the controller.
        - **Span composition**: `spans` is a collection of `storage.Span` objects (cascade on delete), defining the physical makeup of the disk group.
        - **Virtual drive association**: `virtualDrives` links to the virtual drives built from the disk group (unset on delete to avoid forced removal of VDs).
        - **Hot spare modeling**: `dedicatedHotSpares` lists physical drives configured as dedicated hot spares for the disk group.
        - **Controller anchoring**: `storageController` ties the disk group to the controller where it is created (cascade on controller delete).
        - **Inventory context**: `registeredDevice` provides device association for inventory scoping and correlation.

---

# Data Source: intersight_storage_disk_group
DiskGroups represent a group of one or more Spans used to configure virtual drives. They capture disk-group identity and RAID intent and tie together spans, virtual drives, and dedicated hot spares under a specific storage controller.
#### Purpose
Provide a controller-scoped construct for organizing disks (via spans) into RAID-oriented groups that can then produce one or more virtual drives.
#### Key Concepts
- **RAID grouping construct**: `raidType` expresses the RAID level intended for virtual drives created from the disk group.
- **Human-identifiable naming**: `name` identifies the disk group on the controller.
- **Span composition**: `spans` is a collection of `storage.Span` objects (cascade on delete), defining the physical makeup of the disk group.
- **Virtual drive association**: `virtualDrives` links to the virtual drives built from the disk group (unset on delete to avoid forced removal of VDs).
- **Hot spare modeling**: `dedicatedHotSpares` lists physical drives configured as dedicated hot spares for the disk group.
- **Controller anchoring**: `storageController` ties the disk group to the controller where it is created (cascade on controller delete).
- **Inventory context**: `registeredDevice` provides device association for inventory scoping and correlation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_disk_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name to identity this disk group in the controller. 
* `raid_type`:(string) Raid level of the virtual drives in this diskgroup. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
