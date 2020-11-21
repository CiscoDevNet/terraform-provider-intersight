
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_profile"
sidebar_current: "docs-intersight-data-source-hyperflex-cluster-profile"
description: |-
A profile specifying configuration settings for a HyperFlex cluster.
---

# Data Source: intersight_hyperflex._cluster_profile
A profile specifying configuration settings for a HyperFlex cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `data_ip_address`:(string) The storage data IP address for the HyperFlex cluster. 
* `description`:(string) Description of the profile. 
* `hypervisor_type`:(string) The hypervisor type for the HyperFlex cluster.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `mac_address_prefix`:(string) The MAC address prefix in the form of 00:25:B5:XX. 
* `mgmt_ip_address`:(string) The management IP address for the HyperFlex cluster. 
* `mgmt_platform`:(string) The management platform for the HyperFlex cluster.* `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.* `EDGE` - The host servers used in the cluster deployment are standalone severs. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `replication`:(int) The number of copies of each data block written. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `wwxn_prefix`:(string) The WWxN prefix in the form of 20:00:00:25:B5:XX. 
