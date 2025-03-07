---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_cluster_worker_node"
description: |-
        ClusterWorkerNode is an object that tracks the Intersight Appliance's process for adding a worker node.

---

# Data Source: intersight_appliance_cluster_worker_node
ClusterWorkerNode is an object that tracks the Intersight Appliance's process for adding a worker node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_cluster_worker_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds during the software install. 
* `end_time`:(string) End date of the software install. 
* `error_code`:(int) Error code for Intersight Appliance's software install. In case of failure - this code will help decide if software install can be retried. 
* `hostname`:(string) Hostname of the worker node being added. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `remote_dns`:(string) Round robin DNS address, which should be able to resolve the hostnames of all the nodes in the cluster. 
* `reuse_node`:(bool) Indicates if the worker node being added is being reused. 
* `session_id`:(string) Session Moid for the user session. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the software install. UI can modify startTime to re-schedule an install. 
* `status`:(string) Status of the Intersight Appliance's software install.* `NotReady` - Cluster is not ready. Install cannot be triggered.* `Ready` - Cluster is ready. Install can be triggered.* `InProgress` - Install is currently in progress.* `Success` - Install was run and succeeded.* `Fail` - Install was run and failed. 
* `total_nodes`:(int) Total number of nodes in the system. 
* `total_phases`:(int) TotalPhase represents the total number of the install phases for one install. 
 
