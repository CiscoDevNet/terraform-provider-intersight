---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_identity"
description: |-
        VirtualDriveIdentities provide a stable identity mapping for virtual drives created under a Server Profile. They connect the virtual drive “name” (as used in the StoragePolicy/drive definitions) to the actual `storage.VirtualDrive` object instance, while also recording the StoragePolicy and Server Profile context.
        #### Purpose
        Enable reliable correlation and lookup of a profile’s virtual drives across policy intent, deployed inventory, and lifecycle operations by anchoring identity to the Server Profile.
        #### Key Concepts
        - **Name-to-instance correlation:** Bridges a user/policy-facing virtual drive name to the concrete `storage.VirtualDrive` object.
        - **Profile-scoped identity:** Identity is unique per (virtual drive name, server profile), reflecting that drive names are meaningful within a profile context.
        - **Policy provenance:** Captures which StoragePolicy produced the virtual drive, supporting traceability and governance.
        - **Lifecycle-safe references:** Provides stable linkage for querying, drift analysis, or downstream operations that need to refer to the correct virtual drive.

---

# Data Source: intersight_storage_virtual_drive_identity
VirtualDriveIdentities provide a stable identity mapping for virtual drives created under a Server Profile. They connect the virtual drive “name” (as used in the StoragePolicy/drive definitions) to the actual `storage.VirtualDrive` object instance, while also recording the StoragePolicy and Server Profile context.
#### Purpose
Enable reliable correlation and lookup of a profile’s virtual drives across policy intent, deployed inventory, and lifecycle operations by anchoring identity to the Server Profile.
#### Key Concepts
- **Name-to-instance correlation:** Bridges a user/policy-facing virtual drive name to the concrete `storage.VirtualDrive` object.
- **Profile-scoped identity:** Identity is unique per (virtual drive name, server profile), reflecting that drive names are meaningful within a profile context.
- **Policy provenance:** Captures which StoragePolicy produced the virtual drive, supporting traceability and governance.
- **Lifecycle-safe references:** Provides stable linkage for querying, drift analysis, or downstream operations that need to refer to the correct virtual drive.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_virtual_drive_identity.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The VirtualDrive Name which belongs to the Storage VirtualDrive. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
