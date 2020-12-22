---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_profile"
description: |-
  A profile specifying configuration settings for a HyperFlex cluster.
---

# Data Source: intersight_hyperflex_cluster_profile
A profile specifying configuration settings for a HyperFlex cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `data_ip_address`:(string) The storage data IP address for the HyperFlex cluster. 
* `description`:(string) Description of the profile. 
* `host_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
* `hypervisor_control_ip_address`:(string) The hypervisor control virtual IP address for the HyperFlex compute cluster that is used for node/pod management. 
* `hypervisor_type`:(string) The hypervisor type for the HyperFlex cluster.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `mac_address_prefix`:(string) The MAC address prefix in the form of 00:25:B5:XX. 
* `mgmt_ip_address`:(string) The management IP address for the HyperFlex cluster. 
* `mgmt_platform`:(string) The management platform for the HyperFlex cluster.* `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.* `EDGE` - The host servers used in the cluster deployment are standalone severs. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `replication`:(int) The number of copies of each data block written. 
* `storage_type`:(string) The storage type used for the HyperFlex cluster (HyperFlex Storage or 3rd party).* `HyperFlexDp` - The type of storage is HyperFlex Data Platform.* `ThirdParty` - The type of storage is 3rd Party Storage (PureStorage, etc..). 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `wwxn_prefix`:(string) The WWxN prefix in the form of 20:00:00:25:B5:XX. 
