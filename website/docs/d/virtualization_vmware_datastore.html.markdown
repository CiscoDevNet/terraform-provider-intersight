---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_datastore"
description: |-
  The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
---

# Data Source: intersight_virtualization_vmware_datastore
The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accessible`:(bool) Shows if this datastore is accessible. 
* `host_count`:(int) Number of hosts attached to or supported-by this datastore. 
* `identity`:(string) The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference). 
* `maintenance_mode`:(bool) Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `multiple_host_access`:(bool) Indicates if this datastore is connected to multiple hosts. 
* `name`:(string) Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations. 
* `status`:(string) Datastore health status, as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `thin_provisioning_supported`:(bool) Indicates if this datastore supports thin provisioning for files. 
* `type`:(string) A string indicating the type of the datastore (VMFS, NFS, etc).* `Unknown` - The nature of the file system is unknown.* `VMFS` - It is a Virtual Machine Filesystem.* `NFS` - It is a Network File System.* `vSAN` - It is a virtual Storage Area Network file system.* `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server. 
* `un_committed`:(int) Space uncommitted in this datastore in bytes. 
* `url`:(string) The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/'). 
* `vm_count`:(int) Number of virtual machines relying on (using) this datastore. 
