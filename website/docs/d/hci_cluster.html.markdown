---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_cluster"
description: |-
        A HCI cluster reported by Prism Central.

---

# Data Source: intersight_hci_cluster
A HCI cluster reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `build_info_build_type`:(string) The software build type, such as \ release\  or \ debug\  build. 
* `build_info_commit_id`:(string) The software commit id for this build image. 
* `build_info_full_version`:(string) The longer form of software version. It usually includes the commit id. 
* `build_info_short_commit_id`:(string) The short version of the software commit id for this build image. 
* `build_info_version`:(string) The software version from the build. 
* `cluster_arch`:(string) The CPU architecture of the cluster server such as x86_64 and PPC64LE. 
* `cluster_ext_id`:(string) The unique identifier of the cluster. 
* `container_name`:(string) The name of the default container created as part of cluster creation. 
* `cpu_capacity_hz`:(int) The CPU capacity in Hz of the cluster. 
* `cpu_usage_hz`:(int) The CPU usage in Hz of the cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `encryption_in_transit_status`:(bool) Indicate if encryption-in-transit is enabled or not. 
* `external_subnet`:(string) The external subnet of the cluster. 
* `incarnationid`:(string) Cluster incarnation Id, part of payload for cluster update operation only. 
* `inefficient_vm_count`:(int) The number of inefficient VMs in this cluster. 
* `internal_subnet`:(string) The internal subnet of the cluster. 
* `is_lts`:(bool) The LTS status indicates whether the release is categorized as Long-term or not. 
* `key_management_server_type`:(string) The key management server type of the cluster. 
* `masquerading_port`:(int) The masquerading port of the cluster. 
* `memory_capacity_bytes`:(int) The memory capacity in bytes of the cluster. 
* `memory_usage_bytes`:(int) The memory usage in bytes of the cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the reported cluster. 
* `number_of_nodes`:(int) The number of nodes in the cluster. 
* `operation_mode`:(string) The operation mode of the cluster such as NORMAL, READ_ONLY, STAND_ALONE, SWITCH_TO_TWO_NODE, OVERRIDE. 
* `password_remote_login_enabled`:(bool) Indicates whether the password ssh into the cluster is enabled or not. 
* `pc_ext_id`:(string) The unique identifier of the domain manager (Prism Central) instance which manages this cluster. 
* `redundancy_factor`:(int) The redundancy factor of the cluster. 
* `remote_support`:(bool) The remote support status of the cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_capacity_bytes`:(int) The storage capacity in bytes of the cluster. 
* `storage_usage_bytes`:(int) The storage usage in bytes of the cluster. 
* `timezone`:(string) The timezone of the cluster. 
* `upgrade_status`:(string) The upgrade status of a cluster includes the following known values: PENDING, DOWNLOADING, QUEUED, PREUPGRADE, UPGRADING, SUCCEEDED,FAILED, CANCELLED, and SCHEDULED.The upgrade status of a cluster. 
* `vm_count`:(int) The number of VMs running on this cluster. 
 
