---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_ahv_vm"
description: |-
        An AHV VM reported by Prism Central.

---

# Data Source: intersight_hci_ahv_vm
An AHV VM reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_ahv_vm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bios_uuid`:(string) The BIOS UUID of the VM, similar to physical server's serial number. 
* `cluster_ext_id`:(string) The unique identifier of the cluster which owns this VM. 
* `create_time`:(string) The time when this managed object was created. 
* `creation_time`:(string) The time the VM was created. 
* `description`:(string) The description of the VM. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `generation_uuid`:(string) The generation UUID of the VM. 
* `guest_os_name`:(string) The guest OS name of the VM. 
* `hardware_clock_timezone`:(string) VM hardware clock timezone in IANA TZDB format (America/Los_Angeles). Default: UTC. 
* `hypervisor_type`:(string) The hypervisor type of the given VM. It could be AHV, ESX etc. 
* `is_agent_vm`:(bool) Indicates whether the VM is an agent VM or not. When their hostenters maintenance mode, once the normal VMs are evacuated, the agent VMs are powered off. When the host is restored, agent VMs are powered on before the normal VMs are restored. In other words, agent VMs cannot be HA-protected or live migrated. 
* `is_branding_enabled`:(bool) Indicates whether to remove AHV branding from VM firmware tables or not. 
* `is_cpu_hotplug_enabled`:(bool) The CPU hotplug status of the VM. It indicates whether the CPU hotplug feature should be enabled for the VM or not. If enabled, the VM can add or remove vCPUs while the VM is running. 
* `is_cpu_passthrough_enabled`:(bool) The CPU passthrough status of the VM. It Indicates whetherto passthrough the host CPU features to the guest or not. Enabling this will make VM incapable of live migration. 
* `is_cross_cluster_migration_in_progress`:(bool) Indicates whether the VM is in the process of cross cluster migration or not. 
* `is_gpu_console_enabled`:(bool) The GPU console status of the VM. It indicates whether the GPU console should be enabled for the VM or not. If enabled, the VM will have access to the GPU console. 
* `is_live_migrate_capable`:(bool) Indicates whether the VM is live migrate capable or not. If the VM is not live migrate capable, it cannot be live migrated. 
* `is_memory_overcommit_enabled`:(bool) The memory overcommit status of the VM. It indicates whetherthe memory overcommit feature should be enabled for the VM or not.If enabled, parts of the VM memory may reside outside of the hypervisor physical memory. Once enabled, it should be expected that the VM may suffer performance degradation. 
* `is_scsi_controller_enabled`:(bool) The SCSI controller status of the VM. It indicates whether the SCSI controller should be enabled for the VM or not. If enabled, the VM will have access to the SCSI controller. 
* `is_vcpu_hard_pinning_enabled`:(bool) The hard pinning status of the vCPU. It indicates whether the vCPUs should be hard pinned to specific pCPUs or not. 
* `is_vga_console_enabled`:(bool) Indicates whether to enable VGA console for the VM or not. 
* `is_vtpm_enabled`:(bool) Indicates whether the VM has a virtual TPM enabled or not. 
* `machine_type`:(string) The machine type of the VM. Possible values are PC, PSERIES, Q35. 
* `memory_size_bytes`:(int) The memory size in bytes of the VM. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the VM reported by the Prism Central. 
* `node_ext_id`:(string) The unique identifier of the node. 
* `num_cores_per_socket`:(int) The number of cores per socket of the VM. 
* `num_cpu_cores`:(int) The number of CPU cores of the VM. 
* `num_numa_nodes`:(int) Number of NUMA nodes. 0 means NUMA is disabled. 
* `num_sockets`:(int) The number of sockets of the VM. 
* `num_threads_per_core`:(int) The number of threads per core of the VM. 
* `power_state`:(string) The power state of the VM. The possible values are ON, OFF, SUSPENDED (ESXi), PAUSED (AHV), UNDETERMINED.* `UNDETERMINED` - The VM power state is currently unknown.* `OFF` - The VM's power state is powered-off.* `ON` - The VM's power state is powered-on.* `PAUSED` - The VM's power state is paused, applicable only to AHV VM.* `SUSPENDED` - The VM's power state is suspended, applicable only to ESXi VM. 
* `protection_type`:(string) The type of protection applied on a VM. Possible values are UNPROTECTED, PD_PROTECTED, and RULE_PROTECTED.PD_PROTECTED indicates a VM is protected using the Prism Element. RULE_PROTECTED indicates a VM protection using the Prism Central. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_uuid`:(string) The source UUID of the VM. 
* `update_time`:(string) The time the VM was last updated. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
* `vtpm_module_version`:(string) The version of the vTPM module. 
 
