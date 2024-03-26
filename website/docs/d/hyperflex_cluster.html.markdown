---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster"
description: |-
        A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes
        of the cluster.

---

# Data Source: intersight_hyperflex_cluster
A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes
of the cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity_runway`:(int) The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.Default value is math.MaxInt32 to indicate that the capacity runway is \ Unknown\  for a cluster that is not connected or with not sufficient data. 
* `cluster_name`:(string) The name of this HyperFlex cluster. 
* `cluster_purpose`:(string) This can be a Storage or Compute cluster. A storage cluster contains storage nodes that are used to persist data. A compute cluster contains compute nodes that are used for executing business logic.* `Storage` - Cluster of storage nodes used to persist data.* `Compute` - Cluster of compute nodes used to execute business logic.* `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values. 
* `cluster_type`:(int) The storage type of this cluster (All Flash or Hybrid). 
* `cluster_uuid`:(string) The unique identifier for this HyperFlex cluster. 
* `compute_node_count`:(int) The number of compute nodes that belong to this cluster. 
* `converged_node_count`:(int) The number of converged nodes that belong to this cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_type`:(string) The deployment type of the HyperFlex cluster.The cluster can have one of the following configurations:1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,the deployment type is set as 'NA' (not available).* `NA` - The deployment type of the HyperFlex cluster is not available.* `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.* `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.* `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.* `DC-No-FI` - The deployment type of a HyperFlex cluster consisting of 3 or more standalone nodes with the required Datacenter license. 
* `device_id`:(string) The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `drive_type`:(string) The type of the drives used for storage in this cluster.* `NA` - The drive type of the cluster is not available.* `All-Flash` - Indicates that this cluster contains flash drives only.* `Hybrid` - Indicates that this cluster contains both flash and hard disk drives. 
* `encryption_status`:(string) This captures the encryption status for a HyperFlex cluster.Currently it will have the status if HXA-CLU-0020 alarm is raised. In the future it can capture other details. 
* `flt_aggr`:(int) The number of yellow (warning) and red (critical) alarms stored as an aggregate.The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms. 
* `hx_version`:(string) The HyperFlex Data or Application Platform version of this cluster. 
* `hxdp_build_version`:(string) The version and build number of the HyperFlex Data Platform for this cluster.After a cluster upgrade, this version string will be updated on the next inventory cycle to reflectthe newly installed version. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) The version of hypervisor running on this cluster. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `storage_capacity`:(int) The storage capacity in this cluster. 
* `storage_node_count`:(int) The number of storage nodes that belong to this cluster. 
* `storage_utilization`:(float) The storage utilization is computed based on total capacity and current capacity utilization. 
* `upgrade_status`:(string) The upgrade status of the HyperFlex cluster.* `Unknown` - The upgrade status of the HyperFlex cluster could not be determined.* `Ok` - The upgrade of the HyperFlex cluster is complete.* `InProgress` - The upgrade of the HyperFlex cluster is in-progress.* `Failed` - The upgrade of the HyperFlex cluster has failed.* `Waiting` - The upgrade of the HyperFlex cluster is waiting to continue execution. 
* `uplink_speed`:(string) The uplink speed information of the HyperFlex cluster.* `Unknown` - The uplink speed could not be determined. The physical servers are potentially not claimed.* `10G` - The uplink speed is 10G.* `1G` - The uplink speed is 1G. 
* `utilization_percentage`:(float) The storage utilization percentage is computed based on total capacity and current capacity utilization. 
* `utilization_trend_percentage`:(float) The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. 
* `vm_count`:(int) The number of virtual machines present on this cluster. 
* `zone_type`:(string) The type of availability zone used by the cluster. Physical zones are always used in HyperFlexStretched Clusters. Logical zones may be used if a cluster has Logical Availability Zones (LAZ)enabled.* `UNKNOWN` - The type of zone configured on the HyperFlex cluster is not known.* `NOT_CONFIGURED` - The zone type is not configured.* `LOGICAL` - The zone is a logical zone created when the logical availability zones (LAZ) feature is enabled on the HyperFlex cluster.* `PHYSICAL` - The zone is a physical zone configured on a stretched HyperFlex cluster. 
 
