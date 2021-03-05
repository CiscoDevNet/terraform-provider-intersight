---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_cluster"
description: |-
  A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
---

# Data Source: intersight_hyperflex_hxap_cluster
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_cluster.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `datacenter_name`:(string) Datacenter to which the cluster belongs. 
* `failure_reason`:(string) Reason of the failure when cluster is in failed state. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `management_ip_address`:(string) Management IP Address of the cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
* `nr_version`:(string) Product version of HyperFlex compute cluster. 
 