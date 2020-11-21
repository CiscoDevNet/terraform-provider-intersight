
---
layout: "intersight"
page_title: "Intersight: intersight_bios_system_boot_order"
sidebar_current: "docs-intersight-data-source-bios-system-boot-order"
description: |-
Actual Boot Order of the system.
---

# Data Source: intersight_bios._system_boot_order
Actual Boot Order of the system.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `boot_mode`:(string) The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `secure_boot`:(string) Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).* `Disabled` - Set the state of Secure Boot to Disabled.* `Enabled` - Set the state of Secure Boot to Enabled. 
