
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_eula"
sidebar_current: "docs-intersight-data-source-firmware-eula"
description: |-
End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
---

# Data Source: intersight_firmware._eula
End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accepted`:(bool) EULA acceptance status for the account. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `content`:(string) EULA acceptance form content provided by cisco.com. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
