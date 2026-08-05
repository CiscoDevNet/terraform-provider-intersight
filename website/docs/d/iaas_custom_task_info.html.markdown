---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_custom_task_info"
description: |-
        The CustomTaskInfo object offers insights into the execution of custom tasks within UCS Director (UCSD), supporting workflow customization and management.
        #### Purpose
        CustomTaskInfo documents execution details of custom tasks, aiding in workflow personalization and optimization.
        #### Key Concepts
        - **Custom Task Execution:** Tracks execution of custom tasks, supporting process customization and refinement.
        - **Read-Only Access:** Ensures secure access to custom task data without alteration.

---

# Data Source: intersight_iaas_custom_task_info
The CustomTaskInfo object offers insights into the execution of custom tasks within UCS Director (UCSD), supporting workflow customization and management.  
#### Purpose  
CustomTaskInfo documents execution details of custom tasks, aiding in workflow personalization and optimization.  
#### Key Concepts 
- **Custom Task Execution:** Tracks execution of custom tasks, supporting process customization and refinement. 
- **Read-Only Access:** Ensures secure access to custom task data without alteration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_custom_task_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `task_description`:(string) Task decription or Comment of the Custom task. 
* `task_execution_count`:(int) Number of times this task has executed. 
* `task_label`:(string) Task Label in the Workflow. 
* `task_name`:(string) Name of the Custom Task in UCSD. 
 
