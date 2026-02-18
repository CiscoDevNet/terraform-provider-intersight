---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_policy"
description: |-
        The PersistentMemoryPolicy is a reusable policy for configuring Intel Optane Persistent Memory (PMem) modules on supported servers. This provides a comprehensive framework for defining memory operating modes, creating logical namespaces, and managing security.
        #### Purpose
        The purpose of this policy is to automate and standardize the complex configuration of persistent memory. It allows administrators to define how the PMem modules should be partitioned between volatile memory (Memory Mode) and persistent storage (App Direct Mode). It also handles the creation of namespaces and the application of security passphrases, enabling consistent and repeatable deployments of PMem-enabled servers.
        #### Key Concepts
        - **Memory Goals:** The policy uses Goals to define the desired ratio of Memory Mode to App Direct Mode for all PMem modules on a given CPU socket.
        - **Logical Namespaces:** Within the App Direct portion, the policy can define one or more logical namespaces, which appear to the operating system as block or raw (DAX) devices.
        - **Local Security:** It supports enabling encryption on the PMem modules by setting a secure passphrase.
        - **Profile-Based and Reboot Required:** The policy is applied via a Server Profile, and its application requires a server reboot to reconfigure the memory subsystem.
        - **Management Mode:** Allows administrators to choose whether the persistent memory is configured by Intersight or managed directly by the operating system.

---

# Data Source: intersight_memory_persistent_memory_policy
The PersistentMemoryPolicy is a reusable policy for configuring Intel Optane Persistent Memory (PMem) modules on supported servers. This provides a comprehensive framework for defining memory operating modes, creating logical namespaces, and managing security.
#### Purpose
The purpose of this policy is to automate and standardize the complex configuration of persistent memory. It allows administrators to define how the PMem modules should be partitioned between volatile memory (Memory Mode) and persistent storage (App Direct Mode). It also handles the creation of namespaces and the application of security passphrases, enabling consistent and repeatable deployments of PMem-enabled servers. 
#### Key Concepts
- **Memory Goals:** The policy uses "Goals" to define the desired ratio of Memory Mode to App Direct Mode for all PMem modules on a given CPU socket.
- **Logical Namespaces:** Within the App Direct portion, the policy can define one or more logical namespaces, which appear to the operating system as block or raw (DAX) devices.
- **Local Security:** It supports enabling encryption on the PMem modules by setting a secure passphrase.
- **Profile-Based and Reboot Required:** The policy is applied via a Server Profile, and its application requires a server reboot to reconfigure the memory subsystem.
- **Management Mode:** Allows administrators to choose whether the persistent memory is configured by Intersight or managed directly by the operating system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_persistent_memory_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `management_mode`:(string) Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System.* `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy.* `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `retain_namespaces`:(bool) Persistent Memory Namespaces to be retained or not. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
