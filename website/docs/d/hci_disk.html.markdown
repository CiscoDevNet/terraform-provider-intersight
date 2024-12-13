---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_disk"
description: |-
        A disk associated with a node.

---

# Data Source: intersight_hci_disk
A disk associated with a node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_ext_id`:(string) The unique identifier of the cluster. 
* `cluster_name`:(string) The name of the cluster this disk belongs to. 
* `create_time`:(string) The time when this managed object was created. 
* `disk_ext_id`:(string) The unique identifier of the disk. 
* `disk_size_bytes`:(int) The size of the disk in bytes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) The current firmware version of the disk. 
* `has_boot_partitions_only`:(bool) Indicates if the disk is boot only and no disk operation to be run on it. 
* `host_name`:(string) The name of the host the disk is running on. 
* `is_boot_disk`:(bool) Indicate if the disk is a boot disk. 
* `is_data_migrated`:(bool) Indicates whether the disk data is migrated. 
* `is_diagnostic_info_available`:(bool) Indicates whether the diagnostic information is available. 
* `is_error_found_in_logs`:(bool) Indicates whether the error is found in kernel logs. 
* `is_marked_for_removal`:(bool) Indicates whether the disk is marked for removal. 
* `is_mounted`:(bool) Indicates whether the disk is mounted. 
* `is_online`:(bool) Indicates whether the disk is online or offline. 
* `is_password_protected`:(bool) The password protection status of the disk. 
* `is_planned_outage`:(bool) Indicates if diagnostics are running on the Disk. 
* `is_self_encrypting_drive`:(bool) The self-encrypting drive status of the disk. 
* `is_self_managed_nvme`:(bool) Indicates if the NVMe Disk is self managed and no host/CVM reboot is required. 
* `is_spdk_managed`:(bool) Indicates if NVMe device is managed by storage performance development kit (SPDK). 
* `is_suspected_unhealthy`:(bool) Indicates whether the disk is suspected unhealthy. 
* `is_under_diagnosis`:(bool) Indicates whether the disk is under diagnosis. 
* `is_unhealthy`:(bool) Indicates whether the disk is unhealthy. 
* `location`:(int) The location of the disk in the node. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the reported disk. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mount_path`:(string) The mount path of the disk. 
* `node_ext_id`:(string) The unique identifier of the node. 
* `nvme_pcie_path`:(string) The PCIe path of the NVMe disk. 
* `physical_capacity_bytes`:(int) The physical capacity of the disk in bytes. 
* `serial_number`:(string) The serial number of the disk. 
* `service_vm_id`:(string) The unique identifier of the service VM on the node. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the disk such as NORMAL, MARKED_FOR_REMOVAL_BUT_NOT_DETACHABLE,DETACHABLE, DATA_MIGRATION_INITIATED. 
* `storage_pool_ext_id`:(string) The unique identifier of the storage pool. 
* `storage_tier`:(string) The storage tier of the disk such as SSD_PCIE, SSD_SATA, DAS_SATA, CLOUD, SSD_MEM_NVME. 
* `target_firmware_version`:(string) The target firmware version of the disk. 
* `vendor`:(string) The vendor of the reported disk. 
 
