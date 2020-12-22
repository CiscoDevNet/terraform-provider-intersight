---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_policy"
description: |-
  The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.
---

# Data Source: intersight_memory_persistent_memory_policy
The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `management_mode`:(string) Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System.* `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy.* `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `retain_namespaces`:(bool) Persistent Memory Namespaces to be retained or not. 
