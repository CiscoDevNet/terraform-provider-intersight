
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_running_firmware"
sidebar_current: "docs-intersight-data-source-firmware-running-firmware"
description: |-
Running Firmware on an endpoint.
---

# Data Source: intersight_firmware._running_firmware
Running Firmware on an endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `component`:(string) Kind of the firmware - boot-booloader/system/kernel. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `package_version`:(string) Package version which the firmware belongs to. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `type`:(string) The type of the firmware. 
* `version`:(string) The version of the firmware. 
