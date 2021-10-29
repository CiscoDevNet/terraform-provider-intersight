---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_blade"
description: |-
  Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.
---

# Data Source: intersight_compute_blade
Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_blade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_power_state`:(string) The desired power state of the server. 
* `asset_tag`:(string) The user defined asset tag assigned to the server. 
* `available_memory`:(int) The amount of memory available on the server. 
* `bios_post_complete`:(bool) The BIOS POST completion status of the server. 
* `chassis_id`:(string) The id of the chassis that the blade is discovered in. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fault_summary`:(int) The fault summary for the server. 
* `hardware_uuid`:(string) The universally unique hardware identity of the server provided by the manufacturer. 
* `management_mode`:(string) The management mode of the server.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `memory_speed`:(string) The maximum memory speed in MHz available on the server. 
* `mgmt_ip_address`:(string) Management address of the server. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
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
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `scaled_mode`:(string) The mode of the server that determines it is scaled. 
* `serial`:(string) This field identifies the serial of the given component. 
* `service_profile`:(string) The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) The slot number in the chassis that the blade is discovered in. 
* `total_memory`:(int) The total memory available on the server. 
* `user_label`:(string) The user defined label assigned to the server. 
* `uuid`:(string) The universally unique identity of the server. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 