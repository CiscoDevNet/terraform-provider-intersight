---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_gpu_endpoint_descriptor"
description: |-
        The capability.GpuEndpointDescriptor object is a descriptor used to uniquely identify a GPU (Graphics Processing Unit) within the capability catalog. It acts as a definitive inventory record for all supported GPU models.
        #### Purpose
        The main purpose of this object is to maintain a catalog of GPU models, storing key identification information such as vendor, model, pid (Product ID), and partNumber. This catalog enables the system to validate installed GPUs, check their compatibility with specific server platforms, and provide accurate, detailed hardware inventory.
        #### Key Concepts
        - **GPU Cataloging:** Provides a structured repository for managing information about different GPU models.
        - **Unique Identification:** Uniquely defines each GPU type using a combination of vendor, model, and pid.
        - **Platform Compatibility:** The supportedPlatformsPids property links each GPU to the server platforms it is certified for, facilitating automated configuration checks.
        - **Read-Only Reference:** Serves as a source of truth for GPU specifications, used for system validation and inventory reporting.

---

# Data Source: intersight_capability_gpu_endpoint_descriptor
The capability.GpuEndpointDescriptor object is a descriptor used to uniquely identify a GPU (Graphics Processing Unit) within the capability catalog. It acts as a definitive inventory record for all supported GPU models.
#### Purpose
The main purpose of this object is to maintain a catalog of GPU models, storing key identification information such as vendor, model, pid (Product ID), and partNumber. This catalog enables the system to validate installed GPUs, check their compatibility with specific server platforms, and provide accurate, detailed hardware inventory.
#### Key Concepts
- **GPU Cataloging:** Provides a structured repository for managing information about different GPU models.
- **Unique Identification:** Uniquely defines each GPU type using a combination of vendor, model, and pid.
- **Platform Compatibility:** The supportedPlatformsPids property links each GPU to the server platforms it is certified for, facilitating automated configuration checks.
- **Read-Only Reference:** Serves as a source of truth for GPU specifications, used for system validation and inventory reporting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_gpu_endpoint_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description of the gpu. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field is to provide model of the gpu. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `part_number`:(string) This field is to provide partNumber of the gpu. 
* `pid`:(string) This field is to provide pid of the gpu. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field is to provide vendor of the gpu. 
 
