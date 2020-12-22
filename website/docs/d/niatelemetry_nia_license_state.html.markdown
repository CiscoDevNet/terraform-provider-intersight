---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_license_state"
description: |-
  Object available at device scope for license information. This determines the usage of this attribute.
---

# Data Source: intersight_niatelemetry_nia_license_state
Object available at device scope for license information. This determines the usage of this attribute.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `feature_activated`:(string) Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. 
* `license_activated`:(string) Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pid_type`:(string) PID of device being inventoried. This determines the hardware model type of the device. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
