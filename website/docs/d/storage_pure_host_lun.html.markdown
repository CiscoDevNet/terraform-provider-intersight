---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_lun"
description: |-
        The PureHostLun object represents a logical unit number (LUN) in a PureStorage FlashArray, existing only if a volume is connected to a host or host group.
        #### Purpose
        PureHostLun serves to manage the logical connections between volumes and hosts, ensuring organized data access and storage operations within the FlashArray.
        #### Key Concepts
        - **LUN Management:** PureHostLun provides a mechanism for organizing volume connections, facilitating efficient data access.
        - **Connection Type:** It distinguishes between shared connections (via host groups) and direct host connections, optimizing communication pathways.
        - **Secure Access:** PureHostLun employs privilege sets to ensure secure management of LUNs within the storage array.

---

# Data Source: intersight_storage_pure_host_lun
The PureHostLun object represents a logical unit number (LUN) in a PureStorage FlashArray, existing only if a volume is connected to a host or host group.  
#### Purpose  
PureHostLun serves to manage the logical connections between volumes and hosts, ensuring organized data access and storage operations within the FlashArray.  
#### Key Concepts 
- **LUN Management:** PureHostLun provides a mechanism for organizing volume connections, facilitating efficient data access. 
- **Connection Type:** It distinguishes between shared connections (via host groups) and direct host connections, optimizing communication pathways. 
- **Secure Access:** PureHostLun employs privilege sets to ensure secure management of LUNs within the storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_host_lun.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_group_name`:(string) Name of the host group associated with LUN. 
* `host_name`:(string) Name of the host associated with LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared`:(bool) Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 
