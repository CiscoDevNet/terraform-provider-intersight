---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_gpu"
description: |-
        A Gpu associated with a node.

---

# Data Source: intersight_hci_gpu
A Gpu associated with a node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_gpu.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The unique identifier of the cluster. 
* `cluster_name`:(string) The name of the cluster this GPU belongs to. 
* `controller_vm_id`:(string) The unique identifier of the controller VM. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gpu_ext_id`:(string) The unique identifier of the gpu. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_ext_id`:(string) The unique identifier of the node. 
* `number_of_vgpus_allocated`:(int) The number of vGPUs allocated. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
