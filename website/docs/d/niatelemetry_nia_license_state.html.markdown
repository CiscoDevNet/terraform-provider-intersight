
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_license_state"
sidebar_current: "docs-intersight-data-source-niatelemetry-nia-license-state"
description: |-
Object available at device scope for license information. This determines the usage of this attribute.
---

# Data Source: intersight_niatelemetry._nia_license_state
Object available at device scope for license information. This determines the usage of this attribute.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `feature_activated`:(string) Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. 
* `license_activated`:(string) Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid_type`:(string) PID of device being inventoried. This determines the hardware model type of the device. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
