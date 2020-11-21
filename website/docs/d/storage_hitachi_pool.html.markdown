
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_pool"
sidebar_current: "docs-intersight-data-source-storage-hitachi-pool"
description: |-
A pool entity in Hitachi storage array.
---

# Data Source: intersight_storage._hitachi_pool
A pool entity in Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `blocking_mode_blockade`:(string) Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. 
* `blocking_mode_full`:(string) Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `depletion_threshold`:(string) The depletion threshold set for the pool (%). 
* `is_shrinking`:(bool) Whether the pool is shrinking is output. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `monitoring_mode`:(string) Performance monitoring execution mode (monitor mode).* `N/A` - Not available.* `Period mode` - Period mode.* `Continuous mode` - Continuous mode. 
* `monitoring_status`:(string) Status of monitor information. 
* `name`:(string) Human readable name of the pool, limited to 64 characters. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `pool_action_mode`:(string) Execution mode for the pool.* `N/A` - Not available.* `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator).* `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC. 
* `pool_id`:(string) Object ID for the pool. Platforms that use a number should convert it to string. 
* `progress_of_replacing`:(string) Displays the status of the tier relocation processing. 
* `status`:(string) Human readable status of the pool, indicating the current health.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_reserved_capacity`:(int) Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool. 
* `type`:(string) Human readable type of the pool, such as thin, tiered, active-flash, etc. 
* `warning_threshold`:(int) The warning threshold set for the pool (%). 
