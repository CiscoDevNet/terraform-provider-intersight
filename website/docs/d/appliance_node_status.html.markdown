---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_status"
description: |-
        The NodeStatus object represents the health and readiness status of each node in an Intersight Appliance cluster. It aggregates node-level checks, such as readiness, CPU and memory usage, and file system checks, into a high-level operational status, and reports key utilization and node-state signals.
        #### Purpose
        NodeStatus enables operators to monitor and troubleshoot appliance cluster nodes by exposing per-node operational status, Kubernetes readiness state, and core utilization indicators.
        #### Key Concepts
        - **Node Identity Within the Cluster:** `nodeId` is a system-assigned unique identifier for the appliance node, incrementing from 1.
        - **Operational Health Aggregation:** `statusChecks` results determine `operationalStatus`, mapping Critical to Impaired, Warning to AttentionNeeded, and all-OK to Operational.
        - **Kubernetes Readiness State:** `nodeState` indicates whether the node is Down, Preparing, or Ready to host Kubernetes pods.
        - **Resource Utilization Signals:** `cpuUsage` and `memUsage` provide percent utilization for quick capacity and pressure assessment.
        - **File System Health Correlation:** `fileSystemStatuses` links to per-filesystem status objects for deeper storage and file system troubleshooting.
        - **Context Relationships:** `nodeInfo` ties to node identity and configuration information, and `systemStatus` ties the node to the overall appliance system status.

---

# Data Source: intersight_appliance_node_status
The NodeStatus object represents the health and readiness status of each node in an Intersight Appliance cluster. It aggregates node-level checks, such as readiness, CPU and memory usage, and file system checks, into a high-level operational status, and reports key utilization and node-state signals.
#### Purpose
NodeStatus enables operators to monitor and troubleshoot appliance cluster nodes by exposing per-node operational status, Kubernetes readiness state, and core utilization indicators.
#### Key Concepts
- **Node Identity Within the Cluster:** `nodeId` is a system-assigned unique identifier for the appliance node, incrementing from 1.
- **Operational Health Aggregation:** `statusChecks` results determine `operationalStatus`, mapping Critical to Impaired, Warning to AttentionNeeded, and all-OK to Operational.
- **Kubernetes Readiness State:** `nodeState` indicates whether the node is Down, Preparing, or Ready to host Kubernetes pods.
- **Resource Utilization Signals:** `cpuUsage` and `memUsage` provide percent utilization for quick capacity and pressure assessment.
- **File System Health Correlation:** `fileSystemStatuses` links to per-filesystem status objects for deeper storage and file system troubleshooting.
- **Context Relationships:** `nodeInfo` ties to node identity and configuration information, and `systemStatus` ties the node to the overall appliance system status.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_node_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cpu_usage`:(float) Percentage of CPU currently in use. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mem_usage`:(float) Percentage of memory currently in use. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node.The system incrementally assigns identifiers to each node inthe Intersight Appliance cluster starting with a value of 1. 
* `node_state`:(string) State of the node in terms of its readiness to host Kubernetes pods.* `Down` - The node is yet to come up and join as a member of theKubernetes cluster.* `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods.* `Ready` - The node is ready to host Kubernetes pods. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
