---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_virtual_gpu"
description: |-
        Virtual GPU managed by a cluster and reported by Prism Central.

---

# Data Source: intersight_hci_virtual_gpu
Virtual GPU managed by a cluster and reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_virtual_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assignable`:(bool) If the GPU resources is available to be allocated to virtual machines (VMs) within this cluster. 
* `cluster_ext_id`:(string) The unique identifier of the cluster which owns this virtual GPU. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(int) The GPU type of the virtual GPU in an integer format. It is similar to DeviceName which shows the GPU type in string format. 
* `device_name`:(string) The GPU type of the virtual GPU in string format. It is similar to DeviceId which shows the GPU type in integer format. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fraction`:(int) The fraction of the physical GPU assigned. 
* `frame_buffer_size_bytes`:(int) The frame buffer size in bytes of the virtual GPU. 
* `guest_driver_version`:(string) The guest driver version of the virtual GPU. 
* `is_in_use`:(bool) The in use status of the virtual GPU. 
* `max_instances_per_vm`:(int) The maximum instances per VM of the virtual GPU. 
* `max_resolution`:(string) The maximum resolution of the virtual GPU. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `numa_node`:(string) Each GPU in a system may be physically connected to a specific CPU socket or NUMA node. The numaNode value specifies which node the GPU is associated with. In a NUMA system, a computer's memory is divided into multiple nodes. Each node is a combination of processors and a portion of the system's memory. While processors can access memory from all nodes, they have faster access to the memory in their own node compared to memory in other nodes. 
* `number_of_virtual_display_heads`:(int) The number of virtual display heads of the virtual GPU. 
* `sbdf`:(string) The SBDF address of the virtual GPU. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The type of the irtual GPU. Possible values are PASSTHROUGH_GRAPHICS, PASSTHROUGH_COMPUTE, VIRTUAL. 
* `vendor`:(string) The vendor name of the virtual GPU. 
* `virtual_gpu_ext_id`:(string) The unique identifier of the virtual GPU. 
 
