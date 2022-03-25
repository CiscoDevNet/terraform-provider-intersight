---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_machine_snapshot"
description: |-
        The virtual machine snapshot is represented here.

---

# Data Source: intersight_virtualization_vmware_virtual_machine_snapshot
The virtual machine snapshot is represented here.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_virtual_machine_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `creation_time`:(string) Snapshot creation time. Time at which snapshot gets created. 
* `current_snapshot`:(bool) If yes, it determines it is the latest snapshot of the virtual machine. 
* `description`:(string) User provided description of the virtual machine snapshot. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `golden`:(bool) If yes, the virtual machine snapshot cannot be deleted. 
* `identity`:(string) The internally generated identity of the snapshot. This entity is not manipulated by users. It aids in uniquely identifying the snapshot object. For VMware, this is a MOR (managed object reference). 
* `key`:(int) The internally assigned id/key of virtual machine snapshot. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User name provided to identify the snapshot. 
* `predecessor_id`:(int) Predecessor id is the id of the parent snapshot. 
* `quiesced`:(bool) Quiesce pauses all the I/O operations on virtual machine till the snapshot is taken. 
* `ref_value`:(string) Internally assigned MOR reference value. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_size`:(int) Size of the snapshot file created of the virtual machine, stored in bytes. 
 
