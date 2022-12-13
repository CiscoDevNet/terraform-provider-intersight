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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `data_ip_address`:(string) The storage data IP address for the HyperFlex cluster. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
* `hypervisor_control_ip_address`:(string) The hypervisor control virtual IP address for the HyperFlex compute cluster that is used for node/pod management. 
* `hypervisor_type`:(string) The hypervisor type for the HyperFlex cluster.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `is_nic_based`:(bool) The NIC based setup being set/unset determined by inventory. 
* `mac_address_prefix`:(string) The MAC address prefix in the form of 00:25:B5:XX. 
* `mgmt_ip_address`:(string) The management IP address for the HyperFlex cluster. 
* `mgmt_platform`:(string) The management platform for the HyperFlex cluster.* `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.* `EDGE` - The host servers used in the cluster deployment are standalone severs.* `DC-No-FI` - The host servers used in the cluster deployment are standalone servers with the DC Advantage license. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `replication`:(int) The number of copies of each data block written. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_client_ip_address`:(string) The storage data IP address for the HyperFlex cluster. 
* `storage_client_netmask`:(string) The netmask for the Storage client network IP address. 
* `storage_cluster_auxiliary_ip`:(string) The auxiliary storage IP address for the HyperFlex cluster. For two node clusters, this is the IP address of the auxiliary ZK controller. 
* `storage_type`:(string) The storage type used for the HyperFlex cluster (HyperFlex Storage or 3rd party).* `HyperFlexDp` - The type of storage is HyperFlex Data Platform.* `ThirdParty` - The type of storage is 3rd Party Storage (PureStorage, etc..). 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `wwxn_prefix`:(string) The WWxN prefix in the form of 20:00:00:25:B5:XX. 
 
