---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_instance_type"
description: |-
        A policy specifying CPU, Memory and Disk size configuration for a Virtual Machine.

---

# Data Source: intersight_kubernetes_virtual_machine_instance_type
A policy specifying CPU, Memory and Disk size configuration for a Virtual Machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_virtual_machine_instance_type.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cpu`:(int) Number of CPUs allocated to virtual machine. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `disk_size`:(int) Ephemeral disk capacity to be provided with units example - 10Gi. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `memory`:(int) Virtual machine memory defined in mebibytes (MiB). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
