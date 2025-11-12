---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_physical_gpu"
description: |-
        Physical GPU managed by a cluster and reported by Prism Central.

---

# Data Source: intersight_hci_physical_gpu
Physical GPU managed by a cluster and reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_physical_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assignable`:(bool) If the GPU resources is available to be allocated to virtual machines (VMs) within this cluster. 
* `cluster_ext_id`:(string) The unique identifier of the cluster which owns this physical GPU. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(int) The GPU type of the physical GPU in an integer format. It is similar to DeviceNamewhich shows the GPU type in string format. 
* `device_name`:(string) The GPU type of the physical GPU in string format. It is similar to DeviceId which shows the GPU type in integer format. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `frame_buffer_size_bytes`:(int) The frame buffer size in bytes of the physical GPU. 
* `is_in_use`:(bool) The in use status of the physical GPU. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) The mode of the physical GPU. Possible values are - UNUSED, USED_FOR_PASSTHROUGH, USED_FOR_VIRTUAL. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `numa_node`:(string) Each GPU in a system may be physically connected to a specific CPU socket or NUMA node. The numaNode value specifies which node the GPU is associated with. In a NUMA system, a computer's memory is divided into multiple nodes. Each node is a combination of processors and a portion of the system's memory. While processors can access memory from all nodes, they have faster access to the memory in their own node compared to memory in other nodes. 
* `physical_gpu_ext_id`:(string) The unique identifier of the physical GPU. 
* `sbdf`:(string) The SBDF address of the physical GPU. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The type of the physical GPU. Possible values are PASSTHROUGH_GRAPHICS, PASSTHROUGH_COMPUTE, VIRTUAL. 
* `vendor`:(string) The vendor name of the physical GPU. 
 
