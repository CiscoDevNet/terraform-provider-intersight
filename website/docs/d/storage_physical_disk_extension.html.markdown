
---
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk_extension"
sidebar_current: "docs-intersight-data-source-storage-physical-disk-extension"
description: |-
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
---

# Data Source: intersight_storage._physical_disk_extension
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string) The whether disk is bootable or not. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `disk_dn`:(string) The distinguished name of the Physical drive. 
* `disk_id`:(int) The storage Enclosure slotId. 
* `disk_state`:(string) The current drive state of disk. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `health`:(string) The current drive state of disk. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
