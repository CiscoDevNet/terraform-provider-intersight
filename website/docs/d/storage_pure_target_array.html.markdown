---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_target_array"
description: |-
        The PureTargetArray object establishes a connection between FlashArrays for replication purposes, enhancing data availability and recovery.
        #### Purpose
        PureTargetArray facilitates replication between storage arrays, ensuring data redundancy and availability through synchronized data operations.
        #### Key Concepts
        - **Replication Management:** It manages connections for data replication, optimizing data availability and recovery processes.
        - **Array Integration:** Establishes links between source and target arrays, ensuring cohesive replication strategies.
        - **Secure Access:** Utilizes privilege sets to secure and manage replication operations within the storage network.

---

# Data Source: intersight_storage_pure_target_array
The PureTargetArray object establishes a connection between FlashArrays for replication purposes, enhancing data availability and recovery.  
#### Purpose  
PureTargetArray facilitates replication between storage arrays, ensuring data redundancy and availability through synchronized data operations.  
#### Key Concepts  
- **Replication Management:** It manages connections for data replication, optimizing data availability and recovery processes. 
- **Array Integration:** Establishes links between source and target arrays, ensuring cohesive replication strategies.
- **Secure Access:** Utilizes privilege sets to secure and manage replication operations within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_target_array.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Target FlashArray. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
