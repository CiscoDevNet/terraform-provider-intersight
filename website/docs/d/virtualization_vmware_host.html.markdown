
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_host"
sidebar_current: "docs-intersight-data-source-virtualization-vmware-host"
description: |-
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.
---

# Data Source: intersight_virtualization._vmware_host
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `boot_time`:(string) The time when this host booted up. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connection_state`:(string) Indicates if the host is connected to the vCenter. Values are connected, not connected. 
* `hw_power_state`:(string) Is the host Powered-up or Powered-down.* `Unknown` - The entity's power state is unknown.* `PoweredOn` - The entity is powered on.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `Unknown` - The type of the hypervisor is unknown.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - A Cisco HyperFlex Application Platform hypervisor. 
* `identity`:(string) The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). 
* `maintenance_mode`:(bool) Is this host in maintenance mode. Set to true or false. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. 
* `network_adapter_count`:(int) The count of all network adapters attached to this host. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) Serial number of this host (internally generated). 
* `status`:(string) Host health status, as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `storage_adapter_count`:(int) The count of all storage adapters attached to this host. 
* `up_time`:(string) The uptime of the host, stored as Duration (from w3c). 
* `uuid`:(string) Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. 
* `vcenter_host_id`:(string) The identity of this host within vCenter (optional). 
* `vendor`:(string) Commercial vendor details of this hardware. 
