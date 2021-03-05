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
To access the ith object of the results obtained, use `data.intersight_bios_system_boot_order.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `boot_mode`:(string) The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `secure_boot`:(string) Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).* `NotAvailable` - Set the state of Secure Boot to Not Available.* `Disabled` - Set the state of Secure Boot to Disabled.* `Enabled` - Set the state of Secure Boot to Enabled. 
 