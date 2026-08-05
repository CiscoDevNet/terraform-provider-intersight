---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_cluster_operation"
description: |-
        The ClusterOperation object is an inventory managed object that allows users to trigger and track cluster-level operations such as NCC checks, configuration changes, and diagnostics. It acts as a coordination point for workflows that operate on Nutanix clusters.
        #### Purpose
        The ClusterOperation object enables users to initiate operations on a cluster and monitor their progress. Each operation type (e.g., NCC check) is tracked independently with its own status, results, and workflow information.
        #### Key Concepts
        - **Action-based Trigger**: Operations are initiated by setting the appropriate action field (e.g., nutanixClusterCheck.action = Run).
        - **Execution Modes**: Supports both synchronous (wait for workflow trigger) and asynchronous (immediate return) execution via operation's executionMode property.
        - **Status Tracking**: Each operation type maintains its own status in the clusterOpInfos array, preventing status overwrites when multiple operations run concurrently.
        - **Workflow Integration**: Integrates with Intersight workflow engine for operation execution with callbacks for status updates.

---

# Data Source: intersight_hci_cluster_operation
The ClusterOperation object is an inventory managed object that allows users to trigger and track cluster-level operations such as NCC checks, configuration changes, and diagnostics. It acts as a coordination point for workflows that operate on Nutanix clusters.
#### Purpose
The ClusterOperation object enables users to initiate operations on a cluster and monitor their progress. Each operation type (e.g., NCC check) is tracked independently with its own status, results, and workflow information.
#### Key Concepts
- **Action-based Trigger**: Operations are initiated by setting the appropriate action field (e.g., nutanixClusterCheck.action = "Run").  
- **Execution Modes**: Supports both synchronous (wait for workflow trigger) and asynchronous (immediate return) execution via operation's executionMode property.
- **Status Tracking**: Each operation type maintains its own status in the clusterOpInfos array, preventing status overwrites when multiple operations run concurrently.
- **Workflow Integration**: Integrates with Intersight workflow engine for operation execution with callbacks for status updates.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_cluster_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The unique identifier of the cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The display name of the cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
