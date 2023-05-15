---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_system_boot_order"
description: |-
        Actual Boot Order of the system.

---

# Data Source: intersight_bios_system_boot_order
Actual Boot Order of the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_system_boot_order.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `boot_mode`:(string) The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the MBR partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the MBR to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `secure_boot`:(string) Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).* `NotAvailable` - Set the state of Secure Boot to Not Available.* `Disabled` - Set the state of Secure Boot to Disabled.* `Enabled` - Set the state of Secure Boot to Enabled. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
