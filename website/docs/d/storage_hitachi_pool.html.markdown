---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_pool"
description: |-
        The HitachiPool object represents a pool entity within the Hitachi storage array, designed to manage and configure storage pools efficiently. This provides functionalities for monitoring and executing various operations related to storage pool management.
        #### Purpose
        HitachiPool serves as a pivotal element for managing storage pools, supporting the definition of execution modes, monitoring settings, and capacity thresholds.
        #### Key Concepts
        - **Capacity Management:** Handles total reserved capacity and threshold settings, ensuring optimal pool performance and resource allocation.
        - **Monitoring:** Supports performance monitoring execution modes, providing insights into pool operations.
        - **Privilege Sets:** Offers controlled access for pool management tasks with defined privilege sets.
        - **Licensing:** Complies with licensing operations under specified entitlements.

---

# Data Source: intersight_storage_hitachi_pool
The HitachiPool object represents a pool entity within the Hitachi storage array, designed to manage and configure storage pools efficiently. This provides functionalities for monitoring and executing various operations related to storage pool management.
#### Purpose
HitachiPool serves as a pivotal element for managing storage pools, supporting the definition of execution modes, monitoring settings, and capacity thresholds.
#### Key Concepts
- **Capacity Management:** Handles total reserved capacity and threshold settings, ensuring optimal pool performance and resource allocation.
- **Monitoring:** Supports performance monitoring execution modes, providing insights into pool operations.
- **Privilege Sets:** Offers controlled access for pool management tasks with defined privilege sets.
- **Licensing:** Complies with licensing operations under specified entitlements.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `blocking_mode_blockade`:(string) Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. 
* `blocking_mode_full`:(string) Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. 
* `create_time`:(string) The time when this managed object was created. 
* `depletion_threshold`:(string) The depletion threshold set for the pool (%). 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_shrinking`:(bool) Whether the pool is shrinking is output. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `monitoring_mode`:(string) Performance monitoring execution mode (monitor mode).* `N/A` - Performance monitoring is not available.* `Period mode` - Period mode is the default setting. If Period mode is enabled, tier range values and page relocations are determined based solely on the monitoring data from the last complete cycle.* `Continuous mode` - When Continuous mode is enabled, the weighted average efficiency is calculated using the latest monitoring information and the collected monitoring information in the past cycles. Page relocations are determined using this weighted average efficiency. 
* `monitoring_status`:(string) Status of monitor information. 
* `name`:(string) Human readable name of the pool, limited to 64 characters. 
* `pool_action_mode`:(string) Execution mode for the pool.* `N/A` - Execution Mode is not available for the pool.* `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator).* `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC. 
* `pool_id`:(string) Object ID for the pool. Platforms that use a number should convert it to string. 
* `progress_of_replacing`:(string) Displays the status of the tier relocation processing. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Human readable status of the pool, indicating the current health.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_reserved_capacity`:(int) Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool. 
* `type`:(string) Human readable type of the pool, such as thin, tiered, active-flash, etc. 
* `warning_threshold`:(int) The warning threshold set for the pool (%). 
 
