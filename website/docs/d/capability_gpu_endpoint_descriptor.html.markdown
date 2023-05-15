---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_gpu_endpoint_descriptor"
description: |-
        Descriptor that uniquely identifies a cpu.

---

# Data Source: intersight_capability_gpu_endpoint_descriptor
Descriptor that uniquely identifies a cpu.
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
 
