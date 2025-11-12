---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_node"
description: |-
        A HCI node reported by Prism Central.

---

# Data Source: intersight_hci_node
A HCI node reported by Prism Central.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `block_model`:(string) The rackable unit model of the node. 
* `block_serial`:(string) The rackable unit serial number of the node. 
* `boot_time_usecs`:(int) The boot time in microseconds of the node. 
* `cluster_ext_id`:(string) The unique identifier of the cluster. 
* `cluster_name`:(string) The name of the cluster this node belongs to. 
* `controller_vm_id`:(int) The identifier number of the controller VM. 
* `controller_vm_maintanence_mode`:(bool) The maintenance mode status of the controller VM. 
* `controller_vm_nat_port`:(int) The NAT port of the controller VM. 
* `controller_vm_server_uuid`:(string) The Rackable unit UUID of the server. 
* `cpu_capacity_hz`:(int) The CPU capacity in Hz of the node. 
* `cpu_frequency_hz`:(int) The CPU frequency in Hz on the node. 
* `cpu_model`:(string) The CPU model of the node. 
* `cpu_usage_hz`:(int) The CPU usage in Hz of the node. 
* `create_time`:(string) The time when this managed object was created. 
* `default_vhd_container_uuid`:(string) The default VHD container UUID of the node. 
* `default_vhd_location`:(string) The default VHD location of the node. 
* `default_vm_container_uuid`:(string) The default VM container UUID of the node. 
* `default_vm_location`:(string) The default VM location of the node. 
* `disk_count`:(int) The number of disks on the node. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `failover_cluster_fqdn`:(string) The failover cluster FQDN of the node. 
* `failover_cluster_node_status`:(string) The failover cluster node status of the node. 
* `gpu_count`:(int) The number of GPUs on the node. 
* `gpu_driver_version`:(string) The GPU driver version of the node. 
* `has_csr`:(bool) Certificate signing request status of the node. 
* `host_name`:(string) The name of the host the node is running on. 
* `host_type`:(string) The type of the host, e.g. HYPER_CONVERGED, COMPUTE_ONLY, STORAGE_ONLY. 
* `hypervisor_acropolis_connection_state`:(string) The connection state of the hypervisor, e.g. CONNECTED, DISCONNECTED, NOT_AVAILABLE. 
* `hypervisor_number_of_vms`:(int) The number of VMs managed on this node. 
* `hypervisor_state`:(string) The hypervisor state e.g. ACROPOLIS_NORMAL, ENTERING_MAINTENANCE_MODE,ENTERED_MAINTENANCE_MODE, RESERVED_FOR_HA_FAILOVER, ENTERING_MAINTENANCE_MODE_FROM_HA_FAILOVER,RESERVING_FOR_HA_FAILOVER, HA_FAILOVER_SOURCE, HA_FAILOVER_TARGET, HA_HEALING_SOURCE,HA_HEALING_TARGET. 
* `hypervisor_type`:(string) The hypervisor type, e.g. AHV, ESX, HYPERV, XEN, NATIVEHOST etc. 
* `hypervisor_user_name`:(string) The user name of the hypervisor on this node. 
* `hypervisor_version`:(string) The version of the hypervisor on this node. 
* `ipmi_username`:(string) The IPMI user name of the controller. 
* `is_degraded`:(bool) The degraded status of the node. 
* `is_hardware_virtualized`:(bool) The hardware virtualization status of the node. 
* `is_secure_booted`:(bool) The secure boot status of the node. 
* `maintenance_state`:(string) The maintenance state of the node. 
* `memory_capacity_bytes`:(int) The memory capacity in bytes of the node. 
* `memory_size_bytes`:(int) The memory size in bytes of the node. 
* `memory_usage_bytes`:(int) The memory usage in bytes of the node. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_ext_id`:(string) The unique identifier of the node. 
* `node_serial`:(string) The serial number of this node. 
* `node_status`:(string) The status of the node such as NORMAL, TO_BE_REMOVED, OK_TO_BE_REMOVED,NEW_NODE, TO_BE_PREPROTECTED, PREPROTECTED. 
* `number_of_cpu_cores`:(int) The number of CPU cores on the node. 
* `number_of_cpu_sockets`:(int) The number of sockets on the node. 
* `number_of_cpu_threads`:(int) The number of threads on the node. 
* `reboot_pending`:(bool) The reboot pending status of the node. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_capacity_bytes`:(int) The storage capacity in bytes of the node. 
* `storage_usage_bytes`:(int) The storage usage in bytes of the node. 
 
