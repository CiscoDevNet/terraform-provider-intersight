---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_debug_log"
description: |-
        The TaskDebugLog object provides detailed debugging information for task execution within workflows. It serves as a critical tool for isolating and resolving task failures, enhancing workflow reliability.
        #### Purpose
        TaskDebugLog is designed to capture detailed debugging information during task execution, helping with troubleshooting and performance optimization. This provides insights into task execution failures, enabling quick identification of root causes and corrective actions. Additionally, it tracks task retry attempts, offering valuable data for analyzing execution patterns and improving overall reliability.
        #### Key Concepts
        - **Debug Information:** Offers detailed logs of task execution, supporting deep analysis and resolution of workflow issues.
        - **Relationship Mapping:** Links task debug logs to specific workflow instances, ensuring contextual understanding and traceability.

---

# Data Source: intersight_workflow_task_debug_log
The TaskDebugLog object provides detailed debugging information for task execution within workflows. It serves as a critical tool for isolating and resolving task failures, enhancing workflow reliability. 
#### Purpose 
TaskDebugLog is designed to capture detailed debugging information during task execution, helping with troubleshooting and performance optimization. This provides insights into task execution failures, enabling quick identification of root causes and corrective actions. Additionally, it tracks task retry attempts, offering valuable data for analyzing execution patterns and improving overall reliability.
#### Key Concepts  
- **Debug Information:** Offers detailed logs of task execution, supporting deep analysis and resolution of workflow issues. 
- **Relationship Mapping:** Links task debug logs to specific workflow instances, ensuring contextual understanding and traceability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_task_debug_log.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `retry_count`:(int) A counter for number of retries. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `task_inst_id`:(string) The unique identifier for task instance. 
 
