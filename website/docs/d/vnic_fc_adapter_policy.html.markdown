---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_adapter_policy"
description: |-
        The FcAdapterPolicy object manages the configuration and operational parameters of Fibre Channel adapters. It governs how adapters handle traffic, error recovery, and performance tuning.
        #### Purpose
        FcAdapterPolicy centralizes the management of advanced features for Fibre Channel adapters, including error recovery, queue management, and interrupt handling to maximize storage performance.
        #### Key Concepts
        - **Advanced Adapter Settings:** Controls error recovery, queue depths, and I/O throttling.
        - **Integration with vHBAs:** Applied to vHBA objects to govern their runtime behavior.

---

# Data Source: intersight_vnic_fc_adapter_policy
The FcAdapterPolicy object manages the configuration and operational parameters of Fibre Channel adapters. It governs how adapters handle traffic, error recovery, and performance tuning.
#### Purpose
FcAdapterPolicy centralizes the management of advanced features for Fibre Channel adapters, including error recovery, queue management, and interrupt handling to maximize storage performance.
#### Key Concepts
- **Advanced Adapter Settings:** Controls error recovery, queue depths, and I/O throttling.
- **Integration with vHBAs:** Applied to vHBA objects to govern their runtime behavior.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_fc_adapter_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_detection_timeout`:(int) Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. 
* `io_throttle_count`:(int) The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. 
* `lun_count`:(int) The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. Lun Count value can exceed 1024 only for vHBA of type 'FC Initiator' and on servers having supported firmware version. 
* `lun_queue_depth`:(int) The number of commands that the HBA can send and receive in a single transmission per LUN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `resource_allocation_timeout`:(int) Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
