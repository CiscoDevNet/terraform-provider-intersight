---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_system_task_info"
description: |-
        The SystemTaskInfo object provides details about library task executions within UCS Director (UCSD), supporting system oversight and process management.
        #### Purpose
        SystemTaskInfo documents task execution data, aiding in workflow task management and optimization strategies.
        #### Key Concepts
        - **Execution Frequency:** Tracks each library tasks execution count, supporting process refinement and optimization.
        - **Functional Categorization:** Provides information regarding the task category to which the library task belongs in UCSD.
        - **Read-Only Access:** Ensures secure access to system task data without modification.

---

# Data Source: intersight_iaas_system_task_info
The SystemTaskInfo object provides details about library task executions within UCS Director (UCSD), supporting system oversight and process management.  
#### Purpose
SystemTaskInfo documents task execution data, aiding in workflow task management and optimization strategies. 
#### Key Concepts 
- **Execution Frequency:** Tracks each library tasks execution count, supporting process refinement and optimization.
- **Functional Categorization:** Provides information regarding the task category to which the library task belongs in UCSD. 
- **Read-Only Access:** Ensures secure access to system task data without modification.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_system_task_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `task_category`:(string) A functional area to which a task belongs to. 
* `task_execution_count`:(int) Number of times this task has been executed. 
* `task_name`:(string) Task name of the UCSD Task. 
 
