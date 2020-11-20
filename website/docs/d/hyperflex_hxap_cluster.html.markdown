
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_cluster"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-cluster"
description: |-
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
---

# Data Source: intersight_hyperflex._hxap_cluster
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `datacenter_name`:(string) Datacenter to which the cluster belongs. 
* `failure_reason`:(string) Reason of the failure when cluster is in failed state. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `management_ip_address`:(string) Management IP Address of the cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
* `version`:(string) Product version of HyperFlex compute cluster. 
