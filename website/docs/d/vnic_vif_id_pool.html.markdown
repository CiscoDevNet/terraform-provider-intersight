---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_vif_id_pool"
description: |-
        The VifIdPools object manages the pool of unique identifiers (Vif IDs) used for provisioning virtual paths between switches and vNICs/vHBAs.
        #### Purpose
        It ensures that each virtual interface is assigned a unique identifier, which is critical for establishing correct data paths in fabric-attached environments.
        #### Key Concepts
        - **Identifier Allocation:** Tracks available and allocated Vif IDs.
        - **Path Provisioning:** Provides the necessary IDs to set up data paths on the switch.

---

# Data Source: intersight_vnic_vif_id_pool
The VifIdPools object manages the pool of unique identifiers (Vif IDs) used for provisioning virtual paths between switches and vNICs/vHBAs.
#### Purpose
It ensures that each virtual interface is assigned a unique identifier, which is critical for establishing correct data paths in fabric-attached environments.
#### Key Concepts
- **Identifier Allocation:** Tracks available and allocated Vif IDs.
- **Path Provisioning:** Provides the necessary IDs to set up data paths on the switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_vif_id_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_available_id`:(int) Shows the next available channel number ID to be allocated for a vNIC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
