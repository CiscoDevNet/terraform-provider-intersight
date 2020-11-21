
---
layout: "intersight"
page_title: "Intersight: intersight_boot_device_boot_mode"
sidebar_current: "docs-intersight-data-source-boot-device-boot-mode"
description: |-
Boot mode of the devices that BIOS uses to boot them.
---

# Data Source: intersight_boot._device_boot_mode
Boot mode of the devices that BIOS uses to boot them.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `configured_boot_mode`:(string) The user desired BIOS boot mode as configured in the boot policy. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
