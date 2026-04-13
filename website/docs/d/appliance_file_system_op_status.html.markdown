---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_system_op_status"
description: |-
        The FileSystemOpStatus object denotes the operational status of file systems within an Intersight Appliance node. It highlights the capacity and usage metrics critical to maintaining system efficiency.
        #### Purpose
        FileSystemOpStatus ensures the health and functionality of file systems, presenting data on usage and capacity that supports system reliability and performance.
        #### Key Concepts
        - **File System Health:** Determines the operational status through checks that categorize the state as Operational, AttentionNeeded, or Impaired.
        - **Capacity Management:** Provides insights into file system capacity and usage to prevent overutilization and maintain optimal performance.
        - **System Linkages:** Connects with device registration to integrate file system status within broader system health metrics.

---

# Data Source: intersight_appliance_file_system_op_status
The FileSystemOpStatus object denotes the operational status of file systems within an Intersight Appliance node. It highlights the capacity and usage metrics critical to maintaining system efficiency.
#### Purpose
FileSystemOpStatus ensures the health and functionality of file systems, presenting data on usage and capacity that supports system reliability and performance.
#### Key Concepts
- **File System Health:** Determines the operational status through checks that categorize the state as Operational, AttentionNeeded, or Impaired.
- **Capacity Management:** Provides insights into file system capacity and usage to prevent overutilization and maintain optimal performance.
- **System Linkages:** Connects with device registration to integrate file system status within broader system health metrics.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_file_system_op_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(int) Capacity of the file system in bytes. 
* `create_time`:(string) The time when this managed object was created. 
* `disk_order`:(string) Symbolic order identifier of the physical disk from /dev/disk/by-order (e.g., disk1, disk2).This provides a stable, human-readable identifier for the disk that persists across reboots,unlike device names which may change. Currently, the relationship between filesystems to disksis one-to-one and this is unlikely to change in the future. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mountpoint`:(string) Mount point of this file system. 
* `node_hostname`:(string) Hostname of the Intersight Appliance node on which this filesystem is located. 
* `operational_status`:(string) Operational status of the file system.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage`:(float) Percentage of the file system capacity currently in use. 
 
