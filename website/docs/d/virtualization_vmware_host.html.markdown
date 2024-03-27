---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_host"
description: |-
        The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.

---

# Data Source: intersight_virtualization_vmware_host
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_host.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `boot_time`:(string) The time when this host booted up. 
* `connection_state`:(string) Indicates if the host is connected to the vCenter. Values are connected, not connected. 
* `create_time`:(string) The time when this managed object was created. 
* `dc_inv_path`:(string) This field stores the inventory path of a datacenter. Used in host maintenance action. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hw_power_state`:(string) Is the host Powered-up or Powered-down.* `Unknown` - The entity's power state is unknown.* `PoweringOn` - The entity is powering on.* `PoweredOn` - The entity is powered on.* `PoweringOff` - The entity is powering off.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state.* `Rebooting` - The entity reboot is in progress.* `` - The entity's power state is not available. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). 
* `is_ssh_enabled`:(bool) True if SSH is enabled in the host, false otherwise. 
* `maintenance_mode`:(bool) Is this host in maintenance mode. Set to true or false. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. 
* `network_adapter_count`:(int) The count of all network adapters attached to this host. 
* `quarantine_mode`:(bool) Indicates if the host is in quarantine mode. Will be set to True, when in quarantine mode. 
* `serial`:(string) Serial number of this host (internally generated). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Host health status, as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `storage_adapter_count`:(int) The count of all storage adapters attached to this host. 
* `time_zone`:(string) Time zone this host is in. 
* `up_time`:(string) The uptime of the host, stored as Duration (from w3c). 
* `uuid`:(string) Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. 
* `vcenter_host_id`:(string) The identity of this host within vCenter (optional). 
* `vendor`:(string) Commercial vendor details of this hardware. 
 
