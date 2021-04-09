---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_extension"
description: |-
  Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
---

# Data Source: intersight_storage_virtual_drive_extension
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_virtual_drive_extension.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bootable`:(string) The ability to boot from the virtual drive. 
* `container_id`:(int) The container id of the virtual drive. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `drive_state`:(string) The state of the virtual drive. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Virtual drive. 
* `oper_device_id`:(string) The operational device id of the virtual drive. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) The UUID assigned to the virtual drive. 
* `vendor_uuid`:(string) The UUID value of the vendor assigned to the virtual drive. 
* `virtual_drive_dn`:(string) The distinguished name of the virtual drive for which the extended data is provided. 
* `virtual_drive_id`:(string) The Id of the virtual drive. 
 