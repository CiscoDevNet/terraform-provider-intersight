
---
layout: "intersight"
page_title: "Intersight: intersight_bios_boot_device"
sidebar_current: "docs-intersight-data-source-bios-boot-device"
description: |-
Actual boot devices of the system as enumerated by BIOS.
---

# Data Source: intersight_bios._boot_device
Actual boot devices of the system as enumerated by BIOS.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_name`:(string) Name of the Configured Boot Device. 
* `device_type`:(string) Type of the Configured Boot Device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
