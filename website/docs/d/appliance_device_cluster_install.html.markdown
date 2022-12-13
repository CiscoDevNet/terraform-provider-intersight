---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_cluster_install"
description: |-
        DeviceClusterInstall is a singleton that tracks the Intersight Cluster Appliance's install.

---

# Data Source: intersight_appliance_device_cluster_install
DeviceClusterInstall is a singleton that tracks the Intersight Cluster Appliance's install.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_device_cluster_install.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds during the software install. 
* `end_time`:(string) End date of the software install. 
* `error_code`:(int) Error code for Intersight Appliance's software install. In case of failure - this code will help decide if software install can be retried. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the software install. UI can modify startTime to re-schedule an install. 
* `status`:(string) Status of the Intersight Appliance's software install.* `NotReady` - Cluster is not ready. Install cannot be triggered.* `Ready` - Cluster is ready. Install can be triggered.* `InProgress` - Install is currently in progress.* `Success` - Install was run and succeeded.* `Fail` - Install was run and failed. 
* `total_nodes`:(int) Total number of nodes in the system. 
* `total_phases`:(int) TotalPhase represents the total number of the install phases for one install. 
 
