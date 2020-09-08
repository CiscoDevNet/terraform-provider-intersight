
---
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group"
sidebar_current: "docs-intersight-data-source-storage-disk-group"
description: |-
Group of one or more Spans to configure virtual drive.
---

# Data Source: intersight_storage._disk_group
Group of one or more Spans to configure virtual drive.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name to identity this disk group in the controller. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `raid_type`:(string) Raid level of the virtual drives in this diskgroup. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
