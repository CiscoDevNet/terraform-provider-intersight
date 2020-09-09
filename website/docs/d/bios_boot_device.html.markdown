
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_name`:(string) Name of the Configured Boot Device. 
* `device_type`:(string) Type of the Configured Boot Device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
