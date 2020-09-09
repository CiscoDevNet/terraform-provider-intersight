
---
layout: "intersight"
page_title: "Intersight: intersight_boot_precision_policy"
sidebar_current: "docs-intersight-data-source-boot-precision-policy"
description: |-
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
---

# Data Source: intersight_boot._precision_policy
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `configured_boot_mode`:(string) Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.* `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.* `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. 
* `description`:(string) Description of the policy. 
* `enforce_uefi_secure_boot`:(bool) If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
