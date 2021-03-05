---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_container"
description: |-
  A Virtual Disk Drive Container.
---

# Data Source: intersight_storage_virtual_drive_container
A Virtual Disk Drive Container.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_virtual_drive_container.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `container_id`:(int) The identifier for this container. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 