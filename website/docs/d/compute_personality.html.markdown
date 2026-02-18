---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_personality"
description: |-
        The compute.Personality object provides a mechanism to model a server with a defined software personality, separate from its physical hardware identity (PID). This allows for logical grouping or classification of servers based on their role or software stack.
        #### Purpose
        The main purpose of the compute.Personality object is to assign a logical identity to a server. This can be used for organizing servers, applying specific software-defined policies, or identifying servers that belong to a particular solution (e.g., a node in a hyper-converged cluster). This provides a flexible way to manage servers based on their function rather than just their hardware.
        #### Key Concepts
        - **Logical Identity:** Assigns a name and personalityId to a server that is independent of its physical model or serial number.
        - **Flexible Grouping:** Allows servers to be grouped and managed based on their assigned personality.
        - **Extensibility:** The additionalInfo field provides a space for custom metadata related to the personality.
        - **Server Association:** Directly linked to a compute.RackUnit or compute.Blade, applying the logical identity to a specific physical server.

---

# Data Source: intersight_compute_personality
The compute.Personality object provides a mechanism to model a server with a defined software personality, separate from its physical hardware identity (PID). This allows for logical grouping or classification of servers based on their role or software stack.
#### Purpose
The main purpose of the compute.Personality object is to assign a logical identity to a server. This can be used for organizing servers, applying specific software-defined policies, or identifying servers that belong to a particular solution (e.g., a node in a hyper-converged cluster). This provides a flexible way to manage servers based on their function rather than just their hardware.
#### Key Concepts
- **Logical Identity:** Assigns a name and personalityId to a server that is independent of its physical model or serial number.
- **Flexible Grouping:** Allows servers to be grouped and managed based on their assigned personality.
- **Extensibility:** The additionalInfo field provides a space for custom metadata related to the personality.
- **Server Association:** Directly linked to a compute.RackUnit or compute.Blade, applying the logical identity to a specific physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_personality.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `additional_info`:(string) Additional info about the added software personality. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the software personality. 
* `personality_id`:(int) Unique identity of added software personality. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
