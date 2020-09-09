
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `content`:(string) EULA acceptance form content provided by cisco.com. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
