
---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure_disk_slot_ep"
sidebar_current: "docs-intersight-data-source-storage-enclosure-disk-slot-ep"
description: |-
Physical Disk slots on the enclosure.
---

# Data Source: intersight_storage._enclosure_disk_slot_ep
Physical Disk slots on the enclosure.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `drive_path`:(string) This field identifies the zoning configuration applied to  this enclosure slot. 
* `health`:(string) This field identifies the health of the disk inserted in the slot. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `presence`:(string) This field identifies the disk is present in the enclosure slot. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot`:(string) This field represents the slot Id in the storage enclosure. 
