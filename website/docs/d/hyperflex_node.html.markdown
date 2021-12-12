---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node"
description: |-
  A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
---

# Data Source: intersight_hyperflex_node
A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `build_number`:(string) The build number of the hypervisor running on the host. 
* `create_time`:(string) The time when this managed object was created. 
* `display_version`:(string) The user-friendly string representation of the hypervisor version of the host. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name`:(string) The hostname configured for the hypervisor running on the host. 
* `hypervisor`:(string) The type of hypervisor running on the host. 
* `lockdown`:(bool) The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model_number`:(string) The model of the host server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_maintenance_mode`:(string) The status of maintenance mode on the HyperFlex node.* `Unknown` - The maintenance mode status could not be determined.* `InMaintenanceMode` - The node has maintenance mode enabled. The node has been temporarily been relinquished from the cluster to allow for maintenance operations.* `NotInMaintenanceMode` - The node does not have maintenance mode enabled. 
* `node_status`:(string) The operational status of the HyperFlex node.* `Unknown` - The default operational status of a HyperFlex node.* `Invalid` - The status of the node cannot be determined by the storage platform.* `Ready` - The platform node has been acknowledged by the cluster.* `Unpublished` - The node is not yet added to the storage cluster.* `Deleted` - The node has been removed from the cluster.* `Blocked` - The node is blocked from being added to the cluster.* `Blacklisted` - The deprecated value for 'Blocked'. It is included to maintain backwards compatibility with clusters running a HyperFlex Data Platform version older than 5.0(1a).* `Allowed` - The node is allowd to be added to the cluster.* `Whitelisted` - The deprecated value for 'Allowed'. It is included to maintain backwards compatibility with clusters running a HyperFlex Data Platform version older than 5.0(1a).* `InMaintenance` - The node is in maintenance mode. It has been temporarily relinquished from the cluster to allow for maintenance operations such as software upgrades.* `Online` - The node is participating in the storage cluster and is available for storage operations.* `Offline` - The node is part of the storage cluster, but is not available for storage operations. 
* `node_uuid`:(string) The unique identifier of the HyperFlex node. 
* `role`:(string) The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage.* `UNKNOWN` - The role of the HyperFlex cluster node is not known.* `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster.* `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster. 
* `serial_number`:(string) The serial of the host server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the host. Indicates whether the hypervisor is online.* `UNKNOWN` - The host status cannot be determined.* `ONLINE` - The host is online and operational.* `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster.* `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade.* `DEGRADED` - The host is degraded and may not be performing in its full operational capacity. 
* `nr_version`:(string) The version of the hypervisor running on the host. 
 