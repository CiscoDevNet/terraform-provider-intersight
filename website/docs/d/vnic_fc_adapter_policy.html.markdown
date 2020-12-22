---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_adapter_policy"
description: |-
  A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.
---

# Data Source: intersight_vnic_fc_adapter_policy
A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `error_detection_timeout`:(int) Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. 
* `io_throttle_count`:(int) The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. 
* `lun_count`:(int) The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. 
* `lun_queue_depth`:(int) The number of commands that the HBA can send and receive in a single transmission per LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `resource_allocation_timeout`:(int) Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. 
