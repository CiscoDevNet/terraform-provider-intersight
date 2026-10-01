---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_span"
description: |-
        Spans represent groups of physical disks (by slot membership) used as building blocks for configuring storage layouts, specifically as components of a DiskGroup used to create virtual drives.
        #### Purpose
        Model the disk grouping construct that aggregates one or more disks into a “span”, enabling controllers to assemble RAID/disk group configurations from one or more spans.
        #### Key Concepts
        - **Disk grouping unit**: A span is a discrete grouping of disks used as input to higher-level constructs like DiskGroups and virtual drive creation.
        - **Slot-based composition**: `slots` provides a list of drive slot identifiers that belong to the span.
        - **Unique identification**: `spanId` uniquely identifies the span within the context it is reported.
        - **Topology relationships**: `physicalDisks` links the span to the actual `storage.PhysicalDisk` objects that comprise it.
        - **DiskGroup association**: Permission inheritance from `diskGroup` indicates spans are governed/owned through their parent DiskGroup context.

---

# Data Source: intersight_storage_span
Spans represent groups of physical disks (by slot membership) used as building blocks for configuring storage layouts, specifically as components of a DiskGroup used to create virtual drives.
#### Purpose
Model the disk grouping construct that aggregates one or more disks into a “span”, enabling controllers to assemble RAID/disk group configurations from one or more spans.
#### Key Concepts
- **Disk grouping unit**: A span is a discrete grouping of disks used as input to higher-level constructs like DiskGroups and virtual drive creation.
- **Slot-based composition**: `slots` provides a list of drive slot identifiers that belong to the span.
- **Unique identification**: `spanId` uniquely identifies the span within the context it is reported.
- **Topology relationships**: `physicalDisks` links the span to the actual `storage.PhysicalDisk` objects that comprise it.
- **DiskGroup association**: Permission inheritance from `diskGroup` indicates spans are governed/owned through their parent DiskGroup context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_span.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `span_id`:(int) Unique identifier value of this span. 
 
