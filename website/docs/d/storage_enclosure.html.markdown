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
To access the ith object of the results obtained, use `data.intersight_storage_enclosure.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chassis_id`:(int) This represent the chassis-ID that houses the storage enclosure. 
* `description`:(string) This represnets the description for the storage enclosure. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `enclosure_id`:(int) This represnets the Identifier for the storage enclosure. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_slots`:(int) This represent the number of slots present in storage enclosure. 
* `presence`:(string) This represent the availability of storage enclosure. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `server_id`:(int) This represent the server-ID that houses the storage enclosure. 
* `type`:(string) This represent the type of storage enclosure. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 