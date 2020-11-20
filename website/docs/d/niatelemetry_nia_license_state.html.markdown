
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `feature_activated`:(string) Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. 
* `license_activated`:(string) Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pid_type`:(string) PID of device being inventoried. This determines the hardware model type of the device. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
