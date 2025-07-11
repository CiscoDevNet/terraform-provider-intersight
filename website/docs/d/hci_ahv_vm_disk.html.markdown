---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_ahv_vm_disk"
description: |-
        A disk associated with an AHV VM.

---

# Data Source: intersight_hci_ahv_vm_disk
A disk associated with an AHV VM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_ahv_vm_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adsf_volume_group_ext_id`:(string) The extId of the volume group which backs this disk for block access. ADSF (Acropolis Distributed Storage Fabric) is the storage layer of the Nutanix platform that provides a distributed storage system. It abstracts and pools the storage resources across the cluster. 
* `bus_type`:(string) The bus type of the disk. Possible values are 'IDE', 'PCI', 'SCSI', 'SATA', 'SPAPR'. 
* `create_time`:(string) The time when this managed object was created. 
* `disk_ext_id`:(string) The unique identifier of the disk. 
* `disk_size_bytes`:(int) The size of the disk in bytes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `index`:(int) The index of the disk, similar to a slot number on physical machine. 
* `is_flash_mode_enabled`:(bool) Indicates whether the virtual disk is pinned to the hot tier or not. 
* `is_migration_in_progress`:(bool) Indicates if the disk is being migrated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_container_ext_id`:(string) The extId of the storage container which backs this disk. 
* `vm_ext_id`:(string) The unique identifier of the VM. 
 
