---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_precision_policy"
description: |-
  Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
---

# Data Source: intersight_boot_precision_policy
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_precision_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `configured_boot_mode`:(string) Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `description`:(string) Description of the policy. 
* `enforce_uefi_secure_boot`:(bool) If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 