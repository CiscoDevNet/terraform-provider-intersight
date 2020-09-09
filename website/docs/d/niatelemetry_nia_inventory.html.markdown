
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory"
sidebar_current: "docs-intersight-data-source-niatelemetry-nia-inventory"
description: |-
Inventory object available per device scope. This common object holds a device level information.
---

# Data Source: intersight_niatelemetry._nia_inventory
Inventory object available per device scope. This common object holds a device level information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cpu`:(float) CPU usage of device being inventoried. This determines the percentage of CPU resources used. 
* `crash_reset_logs`:(string) Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system. 
* `device_name`:(string) Name of device being inventoried. The name the user assigns to the device is inventoried here. 
* `device_type`:(string) Type of device being inventoried. This determines whether the device is a controller, leaf or spine. 
* `fex_count`:(int) Scale of fabric extendors utilized. 
* `log_in_time`:(string) Last log in time device being inventoried. This determines the last login time on the device. 
* `log_out_time`:(string) Last log out time of device being inventoried. This determines the last logout time on the device. 
* `mac_sec_count`:(int) Number of Macsec configured interfaces on a TOR. 
* `mac_sec_fab_count`:(int) Number of Macsec configured interfaces on a Spine. 
* `macsec_total_count`:(int) Number of total Macsec configured interfaces for all nodes. 
* `memory`:(int) Memory usage of device being inventoried. This determines the percentage of memory resources used. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `route_prefix_count`:(int) Total nuumber of v4 and v6 routes per node. 
* `route_prefix_v4_count`:(int) Scale of v4 routes per node. 
* `route_prefix_v6_count`:(int) Scale of v6 routes per node. 
* `serial`:(string) Serial number of device being invetoried. The serial number is unique per device and will be used as the key. 
* `software_download`:(string) Last software downloaded of device being inventoried. This determines if software download API was used. 
* `version`:(string) Software version of device being inventoried. The various software version values for each device are available on cisco.com. 
