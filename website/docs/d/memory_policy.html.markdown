---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_policy"
description: |-
        The Memory Policy is a reusable policy that allows administrators to manage specific memory-related features on a server.
        #### Purpose
        The purpose of this policy is to provide a simple, centralized way to control supported memory features. Currently, its primary function is to enable or disable the blocklisting of faulty DIMMs, which can help improve system stability by preventing the use of unreliable memory modules.
        #### Key Concepts
        - **Feature Control:** Provides a toggle for specific memory features. As of now, this is focused on DIMM blocklisting.
        - **Reusable and Profile-Based:** As a policy object, it can be attached to multiple Server Profiles to ensure a consistent memory configuration is applied across different servers.
        - **Simplicity:** It offers a straightforward way to manage memory settings without delving into complex BIOS tokens.

---

# Data Source: intersight_memory_policy
The Memory Policy is a reusable policy that allows administrators to manage specific memory-related features on a server.
#### Purpose
The purpose of this policy is to provide a simple, centralized way to control supported memory features. Currently, its primary function is to enable or disable the blocklisting of faulty DIMMs, which can help improve system stability by preventing the use of unreliable memory modules.
#### Key Concepts
- **Feature Control:** Provides a toggle for specific memory features. As of now, this is focused on DIMM blocklisting.
- **Reusable and Profile-Based:** As a policy object, it can be attached to multiple Server Profiles to ensure a consistent memory configuration is applied across different servers.
- **Simplicity:** It offers a straightforward way to manage memory settings without delving into complex BIOS tokens.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_memory_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_dimm_blocklisting`:(bool) Enable DIMM Blocklisting on the server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
