---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_deferred_cleanup_input"
description: |-
        ### Overview
        The cleanup operation input object captures the pending deferred-cleanup operations for a workload
        deployment input.
        #### Purpose
        The cleanup operation input provides the current deferred cleanup that must run after normal deployment once
        workload instances converge on the current deployment input. It allows the system to retain the active
        pending cleanup set and clear it after successful execution.
        #### Key Concepts
        - **Operation Set and Targets:** - Stores cleanup actions as a collection of `CleanupOperationInput` entries, including operation type, blueprint reference, generated object name, and target object references.
        - **Current Pending Cleanup Tracking:** - Works with the deployment input relationship to preserve the active pending cleanup set for the deployment input.
        - **Deployment-Input Gating:** - Deferred cleanup is applied based on workload instances converging on the deployment's current deployment input, without storing post-deployment status on the instances themselves.

---

# Data Source: intersight_workload_deferred_cleanup_input
### Overview
The cleanup operation input object captures the pending deferred-cleanup operations for a workload
deployment input.
#### Purpose
The cleanup operation input provides the current deferred cleanup that must run after normal deployment once
workload instances converge on the current deployment input. It allows the system to retain the active
pending cleanup set and clear it after successful execution.
#### Key Concepts
- **Operation Set and Targets:** - Stores cleanup actions as a collection of `CleanupOperationInput` entries, including operation type, blueprint reference, generated object name, and target object references.
- **Current Pending Cleanup Tracking:** - Works with the deployment input relationship to preserve the active pending cleanup set for the deployment input.
- **Deployment-Input Gating:** - Deferred cleanup is applied based on workload instances converging on the deployment's current deployment input, without storing post-deployment status on the instances themselves.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_deferred_cleanup_input.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
