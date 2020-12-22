---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_vd_member_ep"
description: |-
  Reference to LocalDisk to build up a VirtualDrive.
---

# Data Source: intersight_storage_vd_member_ep
Reference to LocalDisk to build up a VirtualDrive.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_qualifier_reason`:(string) For certain states, indicates the reason why the operState is in that state. 
* `presence`:(string) The presence state of the local disk. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) Role of the disk normal or hot-spare, used by virtual-drive. 
* `span_id`:(string) The span id number of the virtual drive. 
* `vd_member_ep_id`:(int) The local disk slot number as id. 
