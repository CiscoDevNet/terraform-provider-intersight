---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_nvme_raid_configuration"
description: |-
        NvmeRaidConfigurations store NVMe HW-RAID configuration under a server profile for a specific controller. They capture drive-group definitions, virtual drive configurations, dedicated hot spares, and physical disk state changes that are meant to be applied at the endpoint on reboot (used in activation workflows).
        #### Purpose
        Persist the computed NVMe RAID plan per controller for a given Server Profile so the system can apply the correct NVMe HW-RAID configuration during the activation step.
        #### Key Concepts
        - **NVMe HW-RAID activation plan:** Encodes what should be created/changed for NVMe RAID after reboot.
        - **Controller-scoped configuration:** Tied to a specific controller DN/MOID (and series for behavior-specific calculations).
        - **Composite config model:** Includes drive groups (with virtual drives), dedicated hot spares, and disk-state updates.
        - **Profile and policy correlation:** Links back to the StoragePolicy used to generate the plan and the Server Profile where it applies.

---

# Data Source: intersight_storage_nvme_raid_configuration
NvmeRaidConfigurations store NVMe HW-RAID configuration under a server profile for a specific controller. They capture drive-group definitions, virtual drive configurations, dedicated hot spares, and physical disk state changes that are meant to be applied at the endpoint on reboot (used in activation workflows).
#### Purpose
Persist the computed NVMe RAID plan per controller for a given Server Profile so the system can apply the correct NVMe HW-RAID configuration during the activation step.
#### Key Concepts
- **NVMe HW-RAID activation plan:** Encodes what should be created/changed for NVMe RAID after reboot.
- **Controller-scoped configuration:** Tied to a specific controller DN/MOID (and series for behavior-specific calculations).
- **Composite config model:** Includes drive groups (with virtual drives), dedicated hot spares, and disk-state updates.
- **Profile and policy correlation:** Links back to the StoragePolicy used to generate the plan and the Server Profile where it applies.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_nvme_raid_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_dn`:(string) The storage controller Dn Name for which Nvme RAID is created at endpoint. 
* `controller_moid`:(string) The storage controller Moid for which Nvme RAID creation is supported. 
* `controller_series`:(string) Describes series of the installed controller. This will be used in the activation step after reboot to calculate the different parameters w.r.t specific controller series. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
