---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_info"
description: |-
        The NodeInfo object is designed to store the cluster node information of the Intersight Appliance. Created during the appliance's setup, NodeInfo objects are updated periodically to reflect the status of each node within the cluster.
        #### Purpose
        NodeInfo serves as an informative entity that provides details about individual nodes in the Intersight Appliance cluster. It includes crucial information such as node IDs, hostnames, IP configurations, and operational status, thus facilitating node management and monitoring.
        #### Key Concepts
        - **Cluster Node Representation:** Each NodeInfo object corresponds to a node within the Intersight Appliance cluster, providing node-specific data.
        - **Read-Only Data:** Ensures that node information is accessible without the risk of unauthorized changes.
        - **Identity Management:** Utilizes system-assigned unique IDs to maintain node identity and facilitate tracking.
        - **Network Configuration:** Captures both IPv4 and IPv6 configurations, supporting comprehensive network management.

---

# Data Source: intersight_appliance_node_info
The NodeInfo object is designed to store the cluster node information of the Intersight Appliance. Created during the appliance's setup, NodeInfo objects are updated periodically to reflect the status of each node within the cluster.
#### Purpose
NodeInfo serves as an informative entity that provides details about individual nodes in the Intersight Appliance cluster. It includes crucial information such as node IDs, hostnames, IP configurations, and operational status, thus facilitating node management and monitoring.
#### Key Concepts
- **Cluster Node Representation:** Each NodeInfo object corresponds to a node within the Intersight Appliance cluster, providing node-specific data.
- **Read-Only Data:** Ensures that node information is accessible without the risk of unauthorized changes.
- **Identity Management:** Utilizes system-assigned unique IDs to maintain node identity and facilitate tracking.
- **Network Configuration:** Captures both IPv4 and IPv6 configurations, supporting comprehensive network management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_node_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_size`:(string) Current running deployment size for the Intersight Appliance node. Eg. small, medium, large etc. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hostname`:(string) Cluster node's FQDN or IP address. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. 
* `node_type`:(string) The node type of Intersight Virtual Appliance.* `standalone` - Single Node Intersight Virtual Appliance.* `management` - Management node type when Intersight Virtual Appliance is running as management-worker deployment.* `hamanagement` - Management node type when Intersight Virtual Appliance is running as multi node HA deployment.* `metrics` - Metrics node when Intersight Virtual Appliance is running management-metrics node. 
* `operational_status`:(string) Operational status of the Intersight Appliance node.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
