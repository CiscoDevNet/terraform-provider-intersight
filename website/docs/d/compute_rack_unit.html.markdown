---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_rack_unit"
description: |-
        Describes a standalone or FI-attached Rack-mounted server.

---

# Data Source: intersight_compute_rack_unit
Describes a standalone or FI-attached Rack-mounted server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_rack_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_power_state`:(string) The desired power state of the server. 
* `asset_tag`:(string) The user defined asset tag assigned to the server. 
* `available_memory`:(int) The amount of memory available on the server. 
* `bios_post_complete`:(bool) The BIOS POST completion status of the server. 
* `connection_status`:(string) Connectivity Status of RackUnit to Switch - A or B or AB. 
* `cpu_capacity`:(float) Total processing capacity of the server. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dimm_blocklisting`:(string) The DIMM Blocklisting setting configured on the server.* `Unknown` - The configured state is unknown.* `Enabled` - The configured state is enabled.* `Disabled` - The configured state is disabled. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fault_summary`:(int) The fault summary for the server. 
* `front_panel_lock_state`:(string) The actual front panel state of the server.* `None` - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered.* `Lock` - Front Panel of the server is set to Locked state.* `Unlock` - Front Panel of the server is set to Unlocked state. 
* `hardware_uuid`:(string) The universally unique hardware identity of the server provided by the manufacturer. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `kvm_server_state_enabled`:(bool) The KVM server state of the server. 
* `kvm_vendor`:(string) The KVM Vendor for the server. 
* `nr_lifecycle`:(string) The lifecycle state of the server. This will map to the discovery lifecycle as represented in the server Identity object.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.* `SecureEraseInProgress` - Secure Erase is in progress on given physical entity.* `ScrubInProgress` - Scrub is in progress on given physical entity.* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.* `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.* `Removed` - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade.* `Moved` - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade.* `Replaced` - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade.* `MovedAndReplaced` - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. 
* `management_mode`:(string) The management mode of the server.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `memory_speed`:(string) The maximum memory speed in MHz available on the server. 
* `mgmt_ip_address`:(string) Management address of the server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC). When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect along with chassis/server Id. When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller. when this server is configired in IMM mode, the value of this property contains model and chassis/server Id. 
* `num_adaptors`:(int) The total number of network adapters present on the server. 
* `num_cpu_cores`:(int) The total number of CPU cores present on the server. 
* `num_cpu_cores_enabled`:(int) The total number of CPU cores enabled on the server. 
* `num_cpus`:(int) The total number of CPUs present on the server. 
* `num_eth_host_interfaces`:(int) The total number of vNICs which are visible to a host on the server. 
* `num_fc_host_interfaces`:(int) The total number of vHBAs which are visible to a host on the server. 
* `num_threads`:(int) The total number of threads the server is capable of handling. 
* `oper_power_state`:(string) The actual power state of the server. 
* `oper_state`:(string) The operational state of the server. 
* `operability`:(string) The operability of the server. 
* `platform_type`:(string) The platform type of the registered device - whether managed by UCSM or operating in standalone mode. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `server_id`:(int) RackUnit ID that uniquely identifies the server. 
* `service_profile`:(string) The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `topology_scan_status`:(string) To maintain the Topology workflow run status. 
* `total_memory`:(int) The total memory available on the server. 
* `tunneled_kvm`:(bool) The Tunneled vKVM status of the server. 
* `user_label`:(string) The user defined label assigned to the server. 
* `uuid`:(string) The universally unique identity of the server. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
