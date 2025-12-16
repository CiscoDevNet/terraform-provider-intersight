---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_op_status"
description: |-
        The NodeOpStatus object captures the status of individual nodes within the Intersight Appliance. It is essential for understanding node-specific operational metrics and performance indicators.
        #### Purpose
        NodeOpStatus focuses on node-level health and readiness, providing detailed metrics such as CPU and memory usage, file system status, and network connectivity within a cluster.
        #### Key Concepts
        - **Node Health:** Assesses individual nodes' operational status, incorporating checks that reflect Critical, Warning, or Operational states.
        - **Resource Monitoring:** Reports on CPU and memory usage to optimize node performance and readiness for hosting applications.
        - **Cluster Connectivity:** Evaluates network status among nodes, ensuring effective communication and data exchange.

---

# Data Source: intersight_appliance_node_op_status
The NodeOpStatus object captures the status of individual nodes within the Intersight Appliance. It is essential for understanding node-specific operational metrics and performance indicators.
#### Purpose
NodeOpStatus focuses on node-level health and readiness, providing detailed metrics such as CPU and memory usage, file system status, and network connectivity within a cluster.
#### Key Concepts
- **Node Health:** Assesses individual nodes' operational status, incorporating checks that reflect Critical, Warning, or Operational states.
- **Resource Monitoring:** Reports on CPU and memory usage to optimize node performance and readiness for hosting applications.
- **Cluster Connectivity:** Evaluates network status among nodes, ensuring effective communication and data exchange.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_node_op_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cpu_usage`:(float) Percentage of CPU currently in use. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mem_usage`:(float) Percentage of memory currently in use. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_hostname`:(string) Hostname of the Intersight Appliance node. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node.The system incrementally assigns identifiers to each node inthe Intersight Appliance cluster starting with a value of 1. 
* `node_state`:(string) State of the node in terms of its readiness to host Kubernetes pods.* `Down` - The node is yet to come up and join as a member of theKubernetes cluster.* `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods.* `Ready` - The node is ready to host Kubernetes pods. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
