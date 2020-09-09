
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_device_summary"
sidebar_current: "docs-intersight-data-source-equipment-device-summary"
description: |-
Aggregation of properties pertaining to different inventory MOs.
---

# Data Source: intersight_equipment._device_summary
Aggregation of properties pertaining to different inventory MOs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dn`:(string) The distinguished name for the Network Element. 
* `model`:(string) The model information of the Network Element. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) The serial number for the Network Element. 
* `source_object_type`:(string) The source object type of this view MO. 
