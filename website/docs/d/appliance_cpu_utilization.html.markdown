---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_cpu_utilization"
description: |-
        CpuUtilization provides CPU utilization metrics for Intersight Appliance nodes. It tracks the percentage of CPU capacity currently in use, enabling administrators to observe processor load trends and spot potential performance bottlenecks across an appliance cluster.
        #### Purpose
        Expose a read-only, system-owned metric stream for monitoring node CPU load to support operational health monitoring and performance troubleshooting.
        #### Key Concepts
        - **Node-scoped utilization metric**: Extends `appliance.NodeUtilizationMetric`, indicating it is part of a common appliance node metric framework.
        - **Percent-based capacity usage**: Represents CPU consumption as a percentage of total available CPU capacity.
        - **Cluster performance visibility**: Useful for identifying imbalanced load or sustained high CPU conditions across nodes.
        - **System-owned telemetry**: `owner: system` indicates it is produced and maintained by the platform, not configured by users.
        - **Restricted read access**: Available via READ to account and system administrators.

---

# Data Source: intersight_appliance_cpu_utilization
CpuUtilization provides CPU utilization metrics for Intersight Appliance nodes. It tracks the percentage of CPU capacity currently in use, enabling administrators to observe processor load trends and spot potential performance bottlenecks across an appliance cluster.
#### Purpose
Expose a read-only, system-owned metric stream for monitoring node CPU load to support operational health monitoring and performance troubleshooting.
#### Key Concepts
- **Node-scoped utilization metric**: Extends `appliance.NodeUtilizationMetric`, indicating it is part of a common appliance node metric framework.
- **Percent-based capacity usage**: Represents CPU consumption as a percentage of total available CPU capacity.
- **Cluster performance visibility**: Useful for identifying imbalanced load or sustained high CPU conditions across nodes.
- **System-owned telemetry**: `owner: system` indicates it is produced and maintained by the platform, not configured by users.
- **Restricted read access**: Available via READ to account and system administrators.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_cpu_utilization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The end of the measurement window. 
* `meta_field`:(string) Uniquely identifies the source of the metric. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_hostname`:(string) Hostname of the Intersight Appliance node. This human-readable identifier corresponds to the Kubernetes node name used for metric queries and provides an intuitive way for users to correlate resource utilization data with specific physical or virtual machines in the cluster topology. 
* `node_id`:(int) System assigned unique ID of the Intersight Appliance node.The system incrementally assigns identifiers to each node inthe Intersight Appliance cluster starting with a value of 1. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `time`:(string) The start of the measurement window. 
* `utilization`:(float) The percent utilization of the metric. 
 
