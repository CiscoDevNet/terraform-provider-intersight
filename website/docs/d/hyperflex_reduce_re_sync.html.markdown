---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_reduce_re_sync"
description: |-
        The ReduceReSync object details the execution status of reduce re-sync and stale mirror cleanup operations for HyperFlex clusters.
        #### Purpose
        ReduceReSync focuses on the execution of maintenance operations that optimize the performance and integrity of HyperFlex clusters, ensuring effective resource management and system health.
        #### Key Concepts
        - **Operational Maintenance:** Facilitates reduce re-sync and cleanup operations, supporting system optimization and reliability.
        - **Cluster Associations:** Maps operations to cluster associations, ensuring coherent execution and management.
        - **Controlled Access:** Enforces privilege sets to ensure secure and authorized execution of maintenance operations.

---

# Data Source: intersight_hyperflex_reduce_re_sync
The ReduceReSync object details the execution status of reduce re-sync and stale mirror cleanup operations for HyperFlex clusters.
#### Purpose
ReduceReSync focuses on the execution of maintenance operations that optimize the performance and integrity of HyperFlex clusters, ensuring effective resource management and system health.
#### Key Concepts
- **Operational Maintenance:** Facilitates reduce re-sync and cleanup operations, supporting system optimization and reliability.
- **Cluster Associations:** Maps operations to cluster associations, ensuring coherent execution and management.
- **Controlled Access:** Enforces privilege sets to ensure secure and authorized execution of maintenance operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_reduce_re_sync.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `completion_status`:(bool) The HyperFlex reduce re-sync script execution completion status. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execution_output`:(string) The execution output of the script. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
