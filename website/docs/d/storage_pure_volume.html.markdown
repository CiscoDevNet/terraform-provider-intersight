
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume"
sidebar_current: "docs-intersight-data-source-storage-pure-volume"
description: |-
A volume entity in PureStorage FlashArray.
---

# Data Source: intersight_storage._pure_volume
A volume entity in PureStorage FlashArray.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `created`:(string) Creation time of the volume. 
* `description`:(string) Short description about the volume. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) Serial number of the volume. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `source`:(string) Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot. 
