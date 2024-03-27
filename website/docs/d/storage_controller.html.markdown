---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_controller"
description: |-
        Storage Controller present in a server.

---

# Data Source: intersight_storage_controller
Storage Controller present in a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_controller.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connected_sas_expander`:(bool) Storage controller is connected to SAS expander. 
* `controller_flags`:(string) The flags for the storage controller. 
* `controller_id`:(string) The Id of the storage controller. 
* `controller_status`:(string) The current status of controller. 
* `create_time`:(string) The time when this managed object was created. 
* `default_drive_mode`:(string) Auto configuration mode for the newly inserted physical drives. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ecc_bucket_leak_rate`:(int) The ECC bucket leak rate for the Storage Controller in minutes. 
* `foreign_config_present`:(bool) Storage controller has detected disks in foreign config. 
* `hw_revision`:(string) The hardware revision of controller. 
* `hybrid_slots_supported`:(string) U.3 Hybrid Slot Support of the Storage Controller. 
* `interface_type`:(string) Interface types are Sas, Sata, PCH. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `max_volumes_supported`:(int) Maximum virtual drives that can be created on this Storage Controller. 
* `memory_correctable_errors`:(int) The number of memory correctable errors reported by the Storage Controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Storage Controller. 
* `oob_interface_supported`:(string) The CIMC support for out-of-band configuration of controller. 
* `oper_state`:(string) The current operational state of controller. 
* `operability`:(string) Operability state of the storage controller. 
* `pci_addr`:(string) The current pci address of controller. 
* `pci_slot`:(string) The pci slot name for the controller. 
* `persistent_cache_size`:(int) The portion of the cache memory that is persistent, measured in MiB. 
* `pinned_cache_state`:(int) The pinned cache state of the Storage Controller. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `raid_support`:(string) The RAID levels supported by controller. 
* `rebuild_rate`:(string) Logical volume or RAID rebuild rate of Storage Controller. 
* `rebuild_rate_percent`:(int) Logical volume or RAID rebuild rate of Storage Controller. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `self_encrypt_enabled`:(string) Storage controller disk self encryption state. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sub_oem_id`:(string) The Sub OEM identifier of the Storage Controller. 
* `supported_strip_sizes`:(string) The strip sizes in KiB supported by the Storage Controller. 
* `total_cache_size`:(int) The total configured cache memory, measured in MiB. 
* `type`:(string) Controller types are Raid, FlexFlash. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
