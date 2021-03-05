---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk_extension"
description: |-
  Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
---

# Data Source: intersight_storage_physical_disk_extension
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_physical_disk_extension.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string) The whether disk is bootable or not. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `disk_dn`:(string) The distinguished name of the Physical drive. 
* `disk_id`:(int) The storage Enclosure slotId. 
* `disk_state`:(string) The current drive state of disk. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `health`:(string) The current drive state of disk. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 