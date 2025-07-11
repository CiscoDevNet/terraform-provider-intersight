---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_ahv_vm_gpu"
description: |-
        A GPU associated with an AHV VM.

---

# Data Source: intersight_hci_ahv_vm_gpu
A GPU associated with an AHV VM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_ahv_vm_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(int) The device id of the GPU. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fraction`:(int) The fraction of a physical GPU's resources that are allocated to a virtual GPU, e.g. 8 indicate 1/8 of a GPU. 
* `frame_buffer_size_bytes`:(int) The frame buffer size of the GPU. 
* `gpu_ext_id`:(string) The unique identifier of the GPU. 
* `guest_driver_version`:(string) Last determined guest driver version. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) The mode of the GPU such as UNUSED, USED_FOR_PASSTHROUGH, USED_FOR_VIRTUAL. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the GPU allocation such as \ NVIDIA_A16-1B\ . 
* `num_virtual_display_heads`:(int) The number of virtual display heads of the GPU. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor name of the GPU. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
 
