---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_deployment_input"
description: |-
        ### Overview
        The DeploymentInput object captures a versioned snapshot of all input parameters and configuration values for a workload deployment at a specific point in time. It maintains a complete history of how deployment inputs have evolved, enabling traceability and potential rollback scenarios.
        #### Purpose
        DeploymentInput serves as an immutable historical record of deployment configuration states. Each time deployment inputs change, a new DeploymentInput object is created with an incremented generation number. This enables the system to track which workload instances were created with which input configurations and to understand the impact of input changes over time.
        #### Key Concepts
        - **Generation-Based Versioning:** - Each input change creates a new generation with an incremented sequence number, providing a complete audit trail of configuration changes.
        - **Input Composition:** - Combines user-provided inputs at both the workload definition and deployment levels, with deployment-level inputs overriding definition-level inputs where applicable.
        - **Blueprint Integration:** - Resolves all blueprint-specific inputs and merges them into a unified input structure that can be applied to workload instance creation.
        - **Object Generation State:** - Tracks whether managed objects (policies, profiles) were successfully generated from this input configuration, supporting rollback scenarios when object generation fails.

---

# Data Source: intersight_workload_deployment_input
### Overview
The DeploymentInput object captures a versioned snapshot of all input parameters and configuration values for a workload deployment at a specific point in time. It maintains a complete history of how deployment inputs have evolved, enabling traceability and potential rollback scenarios.
#### Purpose
DeploymentInput serves as an immutable historical record of deployment configuration states. Each time deployment inputs change, a new DeploymentInput object is created with an incremented generation number. This enables the system to track which workload instances were created with which input configurations and to understand the impact of input changes over time.
#### Key Concepts
- **Generation-Based Versioning:** - Each input change creates a new generation with an incremented sequence number, providing a complete audit trail of configuration changes.
- **Input Composition:** - Combines user-provided inputs at both the workload definition and deployment levels, with deployment-level inputs overriding definition-level inputs where applicable.
- **Blueprint Integration:** - Resolves all blueprint-specific inputs and merges them into a unified input structure that can be applied to workload instance creation.
- **Object Generation State:** - Tracks whether managed objects (policies, profiles) were successfully generated from this input configuration, supporting rollback scenarios when object generation fails.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_deployment_input.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gen_number`:(int) A sequential number that increments whenever the input changes. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_generation_failed`:(bool) The state of the object generation performed using this input. If object generation fails, then this is set to true when the generated objects are restored to their prior state. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
