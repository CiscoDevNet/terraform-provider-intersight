
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_cluster"
sidebar_current: "docs-intersight-data-source-virtualization-vmware-cluster"
description: |-
A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
---

# Data Source: intersight_virtualization._vmware_cluster
A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `datastore_count`:(int) Count of all datastores associated with this cluster. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `Unknown` - The type of the hypervisor is unknown.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - A Cisco HyperFlex Application Platform hypervisor. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
