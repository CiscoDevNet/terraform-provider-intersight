---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_host_lun"
description: |-
        The HitachiHostLun object represents a host LUN entity within a Hitachi storage array. It facilitates public connections to hosts within host groups, supporting effective data access and communication.
        #### Purpose
        HitachiHostLun is integral to managing host LUNs within Hitachi arrays. It enables seamless connectivity to volumes, ensuring efficient data access and resource allocation.
        #### Key Concepts
        - **Public Connectivity:** Provides shared access to volumes across hosts within a host group, simplifying data access management.
        - **Volume Association:** Links host LUNs to specific volumes, supporting streamlined access and resource allocation.
        - **Host Group Integration:** Connects host LUNs to host groups, ensuring reliable communication pathways.
        - **Security and Control:** Implements settings for port and LUN security, enhancing access control and data protection.

---

# Data Source: intersight_storage_hitachi_host_lun
The HitachiHostLun object represents a host LUN entity within a Hitachi storage array. It facilitates public connections to hosts within host groups, supporting effective data access and communication.
#### Purpose
HitachiHostLun is integral to managing host LUNs within Hitachi arrays. It enables seamless connectivity to volumes, ensuring efficient data access and resource allocation.
#### Key Concepts
- **Public Connectivity:** Provides shared access to volumes across hosts within a host group, simplifying data access management.
- **Volume Association:** Links host LUNs to specific volumes, supporting streamlined access and resource allocation.
- **Host Group Integration:** Connects host LUNs to host groups, ensuring reliable communication pathways.
- **Security and Control:** Implements settings for port and LUN security, enhancing access control and data protection.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_host_lun.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(string) Port ID of the Hitachi host lun. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
 
