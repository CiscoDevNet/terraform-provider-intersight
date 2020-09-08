
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_firmware_summary"
sidebar_current: "docs-intersight-data-source-firmware-firmware-summary"
description: |-
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
---

# Data Source: intersight_firmware._firmware_summary
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bundle_version`:(string) Version details at the bundle level for the each of server. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
