---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_datastore_cluster"
description: |-
  The VMware Datastore cluster entity with its attributes. Datastore cluster is a collection of datastores with shared resources and a shared management interface.
---

# Data Source: intersight_virtualization_vmware_datastore_cluster
The VMware Datastore cluster entity with its attributes. Datastore cluster is a collection of datastores with shared resources and a shared management interface.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_datastore_cluster.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `automation_level`:(string) The global automation level for all virtual machines in this datastore cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `datastore_count`:(int) Number of datastores present in this datastore cluster. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `free_space_threshold`:(int) Minimum level of free space for each datastore that is the threshold for action. 
* `host_count`:(int) Number of hosts attached to or supported-by this datastore cluster. 
* `identity`:(string) The internally generated identity of this datastore cluster. This entity is not manipulated by users. It aids in uniquely identifying the datastore cluster object. For VMware, this is an MOR (managed object reference). 
* `inventory_path`:(string) Inventory path of the Datastore Cluster. 
* `io_latency_threshold`:(int) Minimum I/O latency for each datastore below which I/O load balancing moves are not considered. 
* `io_load_balance_automation_mode`:(string) Storage DRS behavior when it generates recommendations for correcting I/O load imbalance in a datastore cluster. 
* `io_load_imbalance_threshold`:(int) Amount of imbalance that Storage DRS should tolerate. 
* `io_metrics_enabled`:(bool) Is I/O Metrics enabled for this datastore cluster. 
* `min_space_utilization_difference`:(int) Specify how much of an improvement DRS should look for before making a recommendation or performing a migration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Datastore Cluster. 
* `policy_enforcement_automation_mode`:(string) Storage DRS behavior when it generates recommendations for correcting storage and VM policy violations in a datastore cluster. 
* `reservable_percent_threshold`:(int) Storage DRS makes storage migration recommendations if total IOPs reservation of all VMs running on a datastore is higher than the specified threshold. 
* `rule_enforcement_automation_mode`:(string) Storage DRS behavior when it generates recommendations for correcting affinity rule violations in a datastore cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `space_load_balance_automation_mode`:(string) Storage DRS behavior when it generates recommendations for correcting space load imbalance in a datastore cluster. 
* `space_threshold_mode`:(string) Runtime thresholds govern when Storage DRS performs or recommends migrations. 
* `status`:(string) Datastore cluster health status, as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `storage_drs_enabled`:(bool) Is Storage DRS enabled for this datastore cluster. 
* `type`:(string) Type of the Datastore Cluster.* `Unknown` - The nature of the file system is unknown.* `VMFS` - It is a Virtual Machine Filesystem.* `NFS` - It is a Network File System.* `vSAN` - It is a virtual Storage Area Network file system.* `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server. 
* `utilized_space_threshold`:(int) Minimum level of consumed space for each datastore that is the threshold for action. 
* `vm_count`:(int) Number of virtual machines relying on (using) this datastore cluster. 
* `vm_evacuation_automation_mode`:(string) Storage DRS behavior when it generates recommendations for VM evacuations from datastores in a datastore cluster. 
 