
---
layout: "intersight"
page_title: "Intersight: intersight_storage_vd_member_ep"
sidebar_current: "docs-intersight-data-source-storage-vd-member-ep"
description: |-
Reference to LocalDisk to build up a VirtualDrive.
---

# Data Source: intersight_storage._vd_member_ep
Reference to LocalDisk to build up a VirtualDrive.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_qualifier_reason`:(string) For certain states, indicates the reason why the operState is in that state. 
* `presence`:(string) The presence state of the local disk. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) Role of the disk normal or hot-spare, used by virtual-drive. 
* `span_id`:(string) The span id number of the virtual drive. 
* `vd_member_ep_id`:(int) The local disk slot number as id. 
