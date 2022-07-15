---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_nvidia_gpu_product"
description: |-
        Information of a Nvidia GPU product.

---

# Data Source: intersight_kubernetes_nvidia_gpu_product
Information of a Nvidia GPU product.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_nvidia_gpu_product.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Optional description of a product. 
* `device_id`:(int) Device Id of a product, which is unique within a vendor. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `memory_size`:(int) Memory size of a GPU product in GB. 
* `mig_capable`:(bool) True if this Nvidia GPU supports MIG. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Display Name of a product. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor_id`:(int) Vendor Id of a product. Each vendor has a globally unique Id, for example 0x10DE for Nvidia. 
 
