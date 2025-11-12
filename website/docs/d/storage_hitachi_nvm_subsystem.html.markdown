---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_nvm_subsystem"
description: |-
        NVM subsystem entity in Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_nvm_subsystem
NVM subsystem entity in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_nvm_subsystem.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_mode`:(string) Host mode of the NVM subsystem. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) NVM subsystem ID. NVM subsystem is a resource to manage system components in an NVMe-oF connection. 
* `namespace_security_setting`:(string) Namespace security settings. 
* `nvm_subsystem_name`:(string) NVM subsystem name. Can be up to 32 characters long. 
* `nvm_subsystem_nqn`:(string) Subsystem NQN. If the NVM subsystem is virtualized, the NQN of the virtualized NVM subsystem is output. 
* `resource_group_id`:(int) Resource group ID of the resource group to which the NVM subsystem belongs. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `t10pi_mode`:(string) Status of the T10 PI mode of the port. 
* `virtual_nvm_subsystem_id`:(int) The Virtual NVM Subsystem ID property is applicable for use with storage systems in the VSP 5000 series. 
 
