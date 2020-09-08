
---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure"
sidebar_current: "docs-intersight-data-source-storage-enclosure"
description: |-
Storage Enclosure for physical disks.
---

# Data Source: intersight_storage._enclosure
Storage Enclosure for physical disks.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chassis_id`:(int) This represent the chassis-ID that houses the storage enclosure. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) This represnets the description for the storage enclosure. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `enclosure_id`:(int) This represnets the Identifier for the storage enclosure. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_slots`:(int) This represent the number of slots present in storage enclosure. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `presence`:(string) This represent the availability of storage enclosure. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `server_id`:(int) This represent the server-ID that houses the storage enclosure. 
* `type`:(string) This represent the type of storage enclosure. 
* `vendor`:(string) This field identifies the vendor of the given component. 
