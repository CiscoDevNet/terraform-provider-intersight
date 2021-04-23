---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure"
description: |-
  Storage Enclosure for physical disks.
---

# Data Source: intersight_storage_enclosure
Storage Enclosure for physical disks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_enclosure.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chassis_id`:(int) This represent the chassis-ID that houses the storage enclosure. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This represnets the description for the storage enclosure. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enclosure_id`:(int) This represnets the Identifier for the storage enclosure. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_slots`:(int) This represent the number of slots present in storage enclosure. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `server_id`:(int) This represent the server-ID that houses the storage enclosure. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) This represent the type of storage enclosure. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 