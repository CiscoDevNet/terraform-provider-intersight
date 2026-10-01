---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_knox_secure_drive_configuration"
description: |-
        KnoxSecureDriveConfigurations store “secure drives” configuration state under a server profile for a specific storage controller (identified by controller DN/MOID). It captures the controller-scoped drive group configuration and physical disk encryption operations that are intended to be applied after reboot as part of secure drive activation workflows.
        #### Purpose
        Persist controller-specific secure-drive configuration derived from a StoragePolicy and bound to a Server Profile, enabling post-reboot activation to securely configure virtual drives and/or secure JBOD disks.
        #### Key Concepts
        - **Profile-scoped activation artifact:** Represents configuration to be applied at the endpoint during an activation step (commonly after reboot).
        - **Controller targeting:** Uses controller identifiers (DN/MOID/series) to scope which controller the secure-drive operations apply to.
        - **Drive group + disk state intent:** Stores both the intended RAID/drive-group layout and per-disk state/encryption operations.
        - **Policy-to-profile linkage:** Maintains relationships to both the StoragePolicy (intent source) and the Server Profile (application target).

---

# Data Source: intersight_storage_knox_secure_drive_configuration
KnoxSecureDriveConfigurations store “secure drives” configuration state under a server profile for a specific storage controller (identified by controller DN/MOID). It captures the controller-scoped drive group configuration and physical disk encryption operations that are intended to be applied after reboot as part of secure drive activation workflows.
#### Purpose
Persist controller-specific secure-drive configuration derived from a StoragePolicy and bound to a Server Profile, enabling post-reboot activation to securely configure virtual drives and/or secure JBOD disks.
#### Key Concepts
- **Profile-scoped activation artifact:** Represents configuration to be applied at the endpoint during an activation step (commonly after reboot).
- **Controller targeting:** Uses controller identifiers (DN/MOID/series) to scope which controller the secure-drive operations apply to.
- **Drive group + disk state intent:** Stores both the intended RAID/drive-group layout and per-disk state/encryption operations.
- **Policy-to-profile linkage:** Maintains relationships to both the StoragePolicy (intent source) and the Server Profile (application target).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_knox_secure_drive_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_dn`:(string) The storage controller Dn Name for which RAID is created at endpoint. 
* `controller_moid`:(string) The storage controller Moid for which RAID creation is supported. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
