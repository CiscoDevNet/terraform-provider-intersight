
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `dn`:(string) The distinguished name for the Network Element. 
* `model`:(string) The model information of the Network Element. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `serial`:(string) The serial number for the Network Element. 
* `source_object_type`:(string) The source object type of this view MO. 
