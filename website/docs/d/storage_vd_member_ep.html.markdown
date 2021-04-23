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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_vd_member_ep.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_qualifier_reason`:(string) For certain states, indicates the reason why the operState is in that state. 
* `presence`:(string) The presence state of the local disk. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) Role of the disk normal or hot-spare, used by virtual-drive. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `span_id`:(string) The span id number of the virtual drive. 
* `vd_member_ep_id`:(int) The local disk slot number as id. 
 